package types

type SearchBlocksRequestPayload struct {
	Height int
	Size   int
}

type SearchBlockResponsePayload struct {
	ID                    uint   `json:"id"`
	Time                  uint64 `json:"time"`
	Height                uint64 `json:"height"`
	ChainID               uint   `json:"chain_id"`
	ProposerConsAddressID uint   `json:"proposer_cons_address_id"`
	TxIndexed             bool   `json:"tx_indexed"`
	BlockEventsIndexed    bool   `json:"block_events_indexed"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type BlockEvent struct {
	Type       string      `json:"type"`
	Attributes []Attribute `json:"attributes"`
}

type SearchBlockByHeightRequestPayload struct {
	Height          string     `json:"height"`
	BeginBlockEvents []BlockEvent `json:"begin_block_events"`
	EndBlockEvents   []BlockEvent `json:"end_block_events"`
}
