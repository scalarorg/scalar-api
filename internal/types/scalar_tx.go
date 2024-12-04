package types

type GetTransactionsRequestPayload struct {
	Height int64
}

type GetTransactionsResponsePayload struct {
	TxHash string `json:"txhash"`
}

type GetTransactionByHashResponsePayload struct {
	Height    int    `json:"height"`
	Type      string `json:"type"`
	Code      int    `json:"code"`
	Timestamp int64  `json:"timestamp"`
	GasUsed   int64  `json:"gas_used"`
	GasWanted int64  `json:"gas_wanted"`
}

type SearchTransactionsRequestPayload struct {
	From        int    `json:"from"`
	Size        int    `json:"size"`
	MessageID   string `json:"messageId"`
	TxHash      string `json:"txHash"`
	Granularity string `json:"granularity"`
}
