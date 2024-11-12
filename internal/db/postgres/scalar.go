package postgres

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/scalarorg/xchains-api/internal/config"
	"github.com/scalarorg/xchains-api/internal/db/postgres/models"
)

type ScalarClient struct {
	scalarPostgresClient *PostgresClient
}

func NewScalarClient(scalarPostgresClient *PostgresClient) *ScalarClient {
	return &ScalarClient{
		scalarPostgresClient: scalarPostgresClient,
	}
}

func (s *ScalarClient) MigrateTables() error {
	if err := s.MigrateDApps(); err != nil {
		return err
	}
	return nil
}

func (s *ScalarClient) InitDApps(cfg config.InitDAppsConfig) error {
	// Check config chains path and runtime chains path
	if cfg.ConfigChainsPath == "" || cfg.RuntimeChainsPath == "" {
		return fmt.Errorf("config chains path or runtime chains path is not set")
	}

	// Read and parse EVM chains configuration file
	evmConfigPath := filepath.Join(cfg.ConfigChainsPath, cfg.Env, cfg.EvmFileName)
	evmConfigData, err := os.ReadFile(evmConfigPath)
	if err != nil {
		return fmt.Errorf("failed to read EVM config at %s: %w", evmConfigPath, err)
	}

	var evmChains []struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(evmConfigData, &evmChains); err != nil {
		return fmt.Errorf("failed to parse EVM config: %w", err)
	}

	// Process each chain
	for _, chain := range evmChains {
		// Read addresses file for the chain
		addressesPath := filepath.Join(cfg.RuntimeChainsPath, chain.ID, cfg.AddressesFileName)

		// Simply check if file exists
		if _, err := os.Stat(addressesPath); err != nil {
			return fmt.Errorf("addresses file not found: %s", err)
		}

		// Read and parse addresses file
		addressesData, err := os.ReadFile(addressesPath)
		if err != nil {
			return fmt.Errorf("failed to read addresses for chain %s: %w", chain.ID, err)
		}

		var addresses struct {
			MintContract string `json:"mintContract"`
		}
		if err := json.Unmarshal(addressesData, &addresses); err != nil {
			return fmt.Errorf("failed to parse addresses for chain %s: %w", chain.ID, err)
		}

		// Create new DApp entry
		dApp := &models.DApp{
			ChainName:            chain.ID,
			BTCAddressHex:        cfg.ProtocolAddressHex,
			PublicKeyHex:         cfg.ProtocolPubKeyHex,
			SmartContractAddress: addresses.MintContract,
			State:                true,
		}

		if err := s.SaveDApp(dApp); err != nil {
			return fmt.Errorf("failed to save DApp for chain %s: %w", chain.ID, err)
		}
	}

	return nil
}

func (s *ScalarClient) MigrateDApps() error {
	return s.scalarPostgresClient.Db.AutoMigrate(&models.DApp{})
}

func (s *ScalarClient) GetDApps() ([]*models.DApp, error) {
	var dApps []*models.DApp
	if err := s.scalarPostgresClient.Db.Find(&dApps).Error; err != nil {
		return nil, err
	}
	return dApps, nil
}

func (s *ScalarClient) SaveDApp(dApp *models.DApp) error {
	dApps := s.scalarPostgresClient.Db
	return dApps.Create(dApp).Error
}

func (s *ScalarClient) UpdateDApp(dApp *models.DApp) error {
	// Using GORM's db.Model().Where().Updates() pattern
	result := s.scalarPostgresClient.Db.Model(&models.DApp{}).
		Where("id = ?", dApp.ID).
		Updates(map[string]interface{}{
			"chain_name":             dApp.ChainName,
			"btc_address_hex":        dApp.BTCAddressHex,
			"public_key_hex":         dApp.PublicKeyHex,
			"smart_contract_address": dApp.SmartContractAddress,
			"chain_id":               dApp.ChainID,
			"chain_endpoint":         dApp.ChainEndpoint,
			"rpc_url":                dApp.RPCUrl,
			"access_token":           dApp.AccessToken,
			"token_contract_address": dApp.TokenContractAddress,
		})

	return result.Error
}

func (s *ScalarClient) ToggleDApp(ID string) error {
	dApps := s.scalarPostgresClient.Db
	var result models.DApp
	if err := dApps.Where("id = ?", ID).First(&result).Error; err != nil {
		return err
	}
	result.State = !result.State
	return dApps.Save(&result).Error
}

func (s *ScalarClient) DeleteDApp(ID string) error {
	dApps := s.scalarPostgresClient.Db
	return dApps.Delete(&models.DApp{}, ID).Error
}
