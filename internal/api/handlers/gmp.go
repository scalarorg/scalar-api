package handlers

import (
	"net/http"

	"github.com/scalarorg/xchains-api/internal/types"
)

func (h *Handler) GmpSearchGmps(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	gmps, err := h.services.GmpSearchGmps(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}

// No payload
func (h *Handler) GmpGetContracts(request *http.Request) (*Result, *types.Error) {
	result, err := h.services.GmpGetContracts(request.Context())
	if err != nil {
		return nil, err
	}
	return NewResult(result), nil
}

func (h *Handler) GmpGetConfigurations(request *http.Request) (*Result, *types.Error) {
	// FUTURE WORK: Implement pagination
	// paginationKey, err := parsePaginationQuery(request)
	// if err != nil {
	// 	return nil, err
	// }
	gmps, err := h.services.GmpGetConfigurations(request.Context())
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}
