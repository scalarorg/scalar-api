// ## BTCAddressConfig defines the configuration for a BTC address.

package config

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

// TODO: Fix chain following https://github.com/satoshilabs/slips/blob/master/slip-0044.md
var BTCAddressByChainConfigs = map[string]*chaincfg.Params{
	"bitcoin-regtest": &chaincfg.RegressionNetParams,
	"bitcoin-testnet": &chaincfg.TestNet3Params,
	"bitcoin-mainnet": &chaincfg.MainNetParams,
	"bitcoin-simnet":  &chaincfg.SimNetParams,
}

func GetTaprootAddress(chain string, pubkeyX string) (string, error) {
	pubkeyXBytes, err := hex.DecodeString(pubkeyX)
	if err != nil {
		return "", err
	}

	taprootAddress, err := btcutil.NewAddressTaproot(pubkeyXBytes, BTCAddressByChainConfigs[chain])
	if err != nil {
		return "", err
	}

	return taprootAddress.EncodeAddress(), nil
}
