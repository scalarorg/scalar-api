package handlers

import (
	"net/http"

	"github.com/scalarorg/xchains-api/internal/types"
)

// GetScalarGlobalParams godoc
// @Summary Get Scalar global parameters
// @Description Retrieves the global parameters for Scalar, including finality provider details.
// @Produce json
// @Success 200 {object} PublicResponse[services.GlobalParamsPublic] "Global parameters"
// @Router /v1/global-params [get]
func (h *Handler) GetScalarGlobalParams(request *http.Request) (*Result, *types.Error) {
	params := h.services.GetGlobalParamsPublic()
	return NewResult(params), nil
}
