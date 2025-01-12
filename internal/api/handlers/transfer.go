package handlers

import (
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

// TokenSearch godoc
// @Summary Search Transfer transactions
// @Description Search for Transfer transactions with filters
// @Tags token
// @Accept json
// @Produce json
// @Success 200 {object} ListPublicResponse[[]models.TransferDocument] "List of Transfer"
// @Router /v1/token/search [post]
func (h *Handler) TokenSearchTransfers(request *http.Request) (*Result, *types.Error) {
	transferPayload, err := models.ParseQueryOptions(request)
	if err != nil {
		return nil, err
	}
	transfers, total, err := h.services.TokenSearchTransfers(request.Context(), transferPayload)
	if err != nil {
		return nil, err
	}

	return NewListResult(transfers, total), nil
}
