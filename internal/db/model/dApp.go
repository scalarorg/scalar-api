package model

type DAppDocument struct {
	ID                     string `bson:"_id,omitempty"`
	ChainName              string `bson:"chain_name"`
	BTCAddressHex          string `bson:"btc_address_hex"`
	PublicKeyHex           string `bson:"public_key_hex"`
	SmartContractAddress   string `bson:"smart_contract_address"`
	State                  bool   `bson:"state"`
	ChainID                string `bson:"chain_id"`
	ChainEndpoint          string `bson:"chain_endpoint"`
	RPCUrl                 string `bson:"rpc_url"`
	AccessToken            string `bson:"access_token"`
}
