package types

type CreateStakingRequestPayload struct {
	Nework                  string `json:"network"`
	StakerAddress           string `json:"staker_address"` //Btc Staker's address
	StakerPubkey            string `json:"staker_pubkey"`
	PublicKeyHex            string `json:"public_key_hex"`
	DstChainId              string `json:"dst_chain_id"` //UPDATE: New field
	DstSmartContractAddress string `json:"dst_smart_contract_address"`
	DstUserAddress          string `json:"dst_user_address"`
	StakingAmount           uint64 `json:"staking_amount"`
	MintingAmount           uint64 `json:"minting_amount"`
	ServicePubkey           string `json:"service_pubkey"`
	FeeRate                 string `json:"fee_rate"` //Expected fee rate
}
type StakingPsbt struct {
	Psbt string `json:"psbt"`
}

type CreateUnstakingRequestPayload struct {
	Nework          string `json:"network"`
	StakerAddress   string `json:"staker_address"`   //Btc Staker's address
	ReceiverAddress string `json:"receiver_address"` //Btc Receiver's address usually is staker's address
	VaultTxHex      string `json:"vault_tx_hex"`
}
type UnstakingPsbt struct {
	Psbt string `json:"psbt"`
}
