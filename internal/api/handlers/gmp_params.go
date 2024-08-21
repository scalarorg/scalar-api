package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/types"
)

func parseGmpPayload(request *http.Request) (*types.GmpPayload, *types.Error) {
	payload := &types.GmpPayload{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid gmp request payload")
	}
	return payload, nil
}
