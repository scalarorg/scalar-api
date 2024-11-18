package config

import (
	"os"
)

type InitDAppsConfig struct {
	Env                string
	ConfigChainsPath   string
	RuntimeChainsPath  string
	ProtocolAddressHex string
	ProtocolPubKeyHex  string
	EvmFileName        string
	AddressesFileName  string
}

func (cfg *InitDAppsConfig) ReadConfig() error {
	cfg.Env = os.Getenv("ENV")
	cfg.ConfigChainsPath = os.Getenv("CONFIG_CHAINS_PATH")
	cfg.RuntimeChainsPath = os.Getenv("RUNTIME_CHAINS_PATH")
	cfg.ProtocolAddressHex = os.Getenv("PROTOCOL_ADDRESS_HEX")
	cfg.ProtocolPubKeyHex = "0x" + os.Getenv("PROTOCOL_PUB_KEY_HEX")
	cfg.EvmFileName = os.Getenv("EVM_FILE_NAME")
	cfg.AddressesFileName = os.Getenv("ADDRESSES_FILE_NAME")
	return nil
}
