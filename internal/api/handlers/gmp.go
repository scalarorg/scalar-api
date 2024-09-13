package handlers

import (
	"net/http"

	gmpDb "github.com/scalarorg/xchains-api/internal/db/gmp"
	"github.com/scalarorg/xchains-api/internal/types"
)

type EnhancedGMPDocument struct {
	*gmpDb.GMPDocument
	SimplifiedStatus string `json:"simplified_status"`
}

func NewEnhancedGMPDocument(gmp *gmpDb.GMPDocument) *EnhancedGMPDocument {
	return &EnhancedGMPDocument{
		GMPDocument:      gmp,
		SimplifiedStatus: gmp.Status,
	}
}

func NewEnhancedGMPDocuments(gmps []*gmpDb.GMPDocument) []*EnhancedGMPDocument {
	enhancedGMPs := make([]*EnhancedGMPDocument, 0, len(gmps))
	for _, gmp := range gmps {
		enhancedGMPs = append(enhancedGMPs, NewEnhancedGMPDocument(gmp))
	}
	return enhancedGMPs
}

func (h *Handler) GMPStats(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}
func (h *Handler) GMPStatsAVGTimes(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}

func (h *Handler) GMPChart(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}

func (h *Handler) GMPCumulativeVolume(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}

func (h *Handler) GMPTotalVolume(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}

func (h *Handler) GMPTotalFee(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}

func (h *Handler) GMPTotalActiveUsers(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}

func (h *Handler) GMPTopUsers(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}

func (h *Handler) GMPTopITSAssets(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}

func (h *Handler) GMPSearch(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}

	newGmps := NewEnhancedGMPDocuments(gmps)

	return NewResult(newGmps), nil
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

func (h *Handler) GetGMPDataMapping(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}

func (h *Handler) EstimateTimeSpent(request *http.Request) (*Result, *types.Error) {
	gmpPayload, err := parseGmpPayload(request)
	if err != nil {
		return nil, err
	}
	gmps, err := h.services.GMPSearch(request.Context(), gmpPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(gmps), nil
}
