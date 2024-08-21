package model

type DAppDocument struct {
	ID            string `bson:"_id,omitempty"`
	ChainName     string `bson:"chain_name"`
	BTCAddressHex string `bson:"btc_address_hex"`
	PublicKeyHex  string `bson:"public_key_hex"`
	State         bool   `bson:"state"`
}
