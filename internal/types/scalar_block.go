package types

type SearchBlocksRequestPayload struct {
	Height int
	Size   int
}

type SearchBlockResponsePayload struct {
	ID              uint   `json:"id"`
	Hash            string `json:"hash"`
	Time            uint64 `json:"time"`
	Height          uint64 `json:"height"`
	ProposerAddress string `json:"proposer_address"`
	NumTxs          int    `json:"num_txs"`
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
	Height           string       `json:"height"`
	Hash             string       `json:"hash"`
	ProposerAddress  string       `json:"proposer_address"`
	Time             uint64       `json:"time"`
	NumTxs           int          `json:"num_txs"`
	BeginBlockEvents []BlockEvent `json:"begin_block_events"`
	EndBlockEvents   []BlockEvent `json:"end_block_events"`
}
