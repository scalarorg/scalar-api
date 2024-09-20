package vault

type VaultDocument struct {
	ID                           string `json:"id,omitempty"`
	Status                       string `json:"status,omitempty"`
	SimplifiedStatus             string `json:"simplified_status,omitempty"`
	SourceChain                  string `json:"source_chain,omitempty"`
	DestinationChain             string `json:"destination_chain,omitempty"`
	DestinationSmartContractAddress string `json:"destination_smart_contract_address,omitempty"`
	CreatedAt                    uint64 `json:"created_at,omitempty"`
	UpdatedAt                    uint64 `json:"updated_at,omitempty"`
	SourceTxHash                 string `json:"source_tx_hash,omitempty"`
	SourceTxHex                  string `json:"source_tx_hex,omitempty"`
	StakerPubkey                 string `json:"staker_pubkey,omitempty"`
	Amount                       string `json:"amount,omitempty"`
}
