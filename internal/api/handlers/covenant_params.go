package handlers

import (
	"net/http"

	"github.com/scalarorg/xchains-api/internal/types"
)

func (h *Handler) GetCovenantParams(request *http.Request) (*Result, *types.Error) {
	covenantParams := types.GetCovenantParamsVar()
	return NewResult(covenantParams), nil
}
