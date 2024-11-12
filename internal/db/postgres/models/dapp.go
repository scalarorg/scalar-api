package models

type DApp struct {
	ID                   string `gorm:"primaryKey"`
	ChainName            string `gorm:"column:chain_name;not null"`
	BTCAddressHex        string `gorm:"column:btc_address_hex"`
	PublicKeyHex         string `gorm:"column:public_key_hex"`
	SmartContractAddress string `gorm:"column:smart_contract_address"`
	State                bool   `gorm:"column:state"`
	ChainID              string `gorm:"column:chain_id"`
	ChainEndpoint        string `gorm:"column:chain_endpoint"`
	RPCUrl               string `gorm:"column:rpc_url"`
	AccessToken          string `gorm:"column:access_token"`
	TokenContractAddress string `gorm:"column:token_contract_address"`
}
