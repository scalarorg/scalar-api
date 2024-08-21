package handlers

import (
	"net/http"

	"github.com/babylonchain/staking-api-service/internal/types"
)

func (h *Handler) GetGMPs(request *http.Request) (*Result, *types.Error) {
	// FUTURE WORK: Implement pagination
	// paginationKey, err := parsePaginationQuery(request)
	// if err != nil {
	// 	return nil, err
	// }
	gmps, err := h.services.GetGMPs(request.Context())
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}
