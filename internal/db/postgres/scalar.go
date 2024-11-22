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
	return s.scalarPostgresClient.Db.AutoMigrate(
		&models.Custodial{},
		&models.DApp{},
		&models.CustodialGroup{},
		&RelayData{},
		&ContractCall{},
		&ContractCallWithToken{},
		&ContractCallApproved{},
		&ContractCallWithTokenApproved{},
		&CommandExecuted{},
	)
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

func (s *ScalarClient) GetDApps() ([]*models.DApp, error) {
	var dApps []*models.DApp
	if err := s.scalarPostgresClient.Db.Preload("CustodialGroup").Find(&dApps).Error; err != nil {
		return nil, err
	}
	return dApps, nil
}

func (s *ScalarClient) SaveDApp(dApp *models.DApp) error {
	dApps := s.scalarPostgresClient.Db
	return dApps.Create(dApp).Error
}

func (s *ScalarClient) UpdateDApp(dApp *models.DApp) error {
	// First find the existing DApp
	existingDApp := &models.DApp{}
	if err := s.scalarPostgresClient.Db.First(existingDApp, dApp.ID).Error; err != nil {
		return err
	}

	// Update the DApp with all fields including associations
	result := s.scalarPostgresClient.Db.Model(existingDApp).
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
			"custodial_group_id":     dApp.CustodialGroupID,
		})

	if result.Error != nil {
		return result.Error
	}

	// Update the CustodialGroup association
	if err := s.scalarPostgresClient.Db.Model(existingDApp).Association("CustodialGroup").Replace(dApp.CustodialGroup); err != nil {
		return err
	}

	return nil
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
	return dApps.Where("id = ?", ID).Delete(&models.DApp{}).Error
}
