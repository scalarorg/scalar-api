package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/scalarorg/xchains-api/internal/types"
)

func parseSearchBlocksPayload(request *http.Request) (*types.SearchBlocksRequestPayload, *types.Error) {
	payload := &types.SearchBlocksRequestPayload{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid search blocks request payload")
	}
	return payload, nil
}

func (h *Handler) SearchBlocks(request *http.Request) (*Result, *types.Error) {
	searchPayload, err := parseSearchBlocksPayload(request)
	if err != nil {
		return nil, err
	}
	blocks, err := h.services.SearchBlocks(request.Context(), searchPayload)
	if err != nil {
		return nil, err
	}

	return NewResult(blocks), nil
}

func (h *Handler) SearchBlockByHeight(request *http.Request) (*Result, *types.Error) {
	height := chi.URLParam(request, "height")
	if height == "" {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "height is required")
	}
	block, err := h.services.SearchBlockByHeight(request.Context(), height)
	if err != nil {
		return nil, err
	}
	return NewResult(block), nil
}
