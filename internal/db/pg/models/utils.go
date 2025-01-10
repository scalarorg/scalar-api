package models

import (
	"encoding/json"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/types"
)

type Options struct {
	Size       int      `json:"size,omitempty"`
	Offset     int      `json:"offset,omitempty"`
	EventId    string   `json:"event_id,omitempty"`
	TxHash     string   `json:"tx_hash,omitempty"`
	EventType  string   `json:"event_type,omitempty"`
	EventTypes []string `json:"event_types,omitempty"`
	// StakerPubkey string
}

func ParseQueryOptions(request *http.Request) (*Options, *types.Error) {
	payload := &Options{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid query options")
	}
	return payload, nil
}
