package pg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/scalarorg/xchains-api/internal/config"
	"github.com/scalarorg/xchains-api/internal/db/pg/models"
)

func (s *PostgresClient) InitDApps(cfg config.InitDAppsConfig) error {
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

func (s *PostgresClient) GetDApps() ([]*models.DApp, error) {
	var dApps []*models.DApp
	if err := s.DB.Preload("CustodialGroup").Preload("CustodialGroup.Custodials").Find(&dApps).Error; err != nil {
		return nil, err
	}
	return dApps, nil
}

func (s *PostgresClient) SaveDApp(dApp *models.DApp) error {
	return s.DB.Create(dApp).Error
}

func (s *PostgresClient) UpdateDApp(dApp *models.DApp) error {
	// First find the existing DApp
	existingDApp := &models.DApp{}
	if err := s.DB.First(existingDApp, dApp.ID).Error; err != nil {
		return err
	}

	// Update the DApp with all fields including associations
	result := s.DB.Model(existingDApp).
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
	if err := s.DB.Model(existingDApp).Association("CustodialGroup").Replace(dApp.CustodialGroup); err != nil {
		return err
	}

	return nil
}

func (s *PostgresClient) ToggleDApp(ID string) error {
	var result models.DApp
	if err := s.DB.Where("id = ?", ID).First(&result).Error; err != nil {
		return err
	}
	result.State = !result.State
	return s.DB.Save(&result).Error
}

func (s *PostgresClient) DeleteDApp(ID string) error {
	return s.DB.Where("id = ?", ID).Delete(&models.DApp{}).Error
}
