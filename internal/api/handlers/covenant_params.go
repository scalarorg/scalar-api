package handlers

import (
	"net/http"

	"github.com/scalarorg/xchains-api/internal/types"
)

func (h *Handler) GetCovenantParams(request *http.Request) (*Result, *types.Error) {
	return NewResult(types.CovenantParams), nil
}
