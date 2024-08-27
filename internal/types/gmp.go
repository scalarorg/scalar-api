package types

type GmpPayload struct {
	From        int    `json:"from"`
	Size        int    `json:"size"`
	MessageID   string `json:"messageId"`
	TxHash      string `json:"txHash"`
	Granularity string `json:"granularity"`
}
