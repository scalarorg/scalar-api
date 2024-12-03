package types

type SearchBlocksRequestPayload struct {
	Height int
	Size   int
}

type SearchBlocksResponsePayload struct {
	ID                    uint   `json:"id"`
	Time                  uint64 `json:"time"`
	Height                uint64 `json:"height"`
	ChainID               uint   `json:"chain_id"`
	ProposerConsAddressID uint   `json:"proposer_cons_address_id"`
	TxIndexed             bool   `json:"tx_indexed"`
	BlockEventsIndexed    bool   `json:"block_events_indexed"`
}
