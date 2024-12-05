package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/scalarorg/xchains-api/internal/types"
)

func NewScalarTxResult[T any](data T, total int) *Result {
	res := &ScalarTxPublicResponse[T]{Data: data, Total: total}
	return &Result{Data: res, Status: http.StatusOK}
}

type ScalarTxPublicResponse[T any] struct {
	Data  T   `json:"data"`
	Total int `json:"total,omitempty"`
}

func parseGetTransactionsPayload(request *http.Request) (*types.GetTransactionsRequestPayload, *types.Error) {
	payload := &types.GetTransactionsRequestPayload{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid get transactions request payload")
	}
	return payload, nil
}

func parseSearchTransactionsPayload(request *http.Request) (*types.SearchTransactionsRequestPayload, *types.Error) {
	payload := &types.SearchTransactionsRequestPayload{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid search transactions request payload")
	}
	return payload, nil
}

func (h *Handler) GetTransactions(request *http.Request) (*Result, *types.Error) {
	payload, err := parseGetTransactionsPayload(request)
	if err != nil {
		return nil, err
	}
	transactions, err := h.services.GetTransactions(request.Context(), payload)
	if err != nil {
		return nil, err
	}
	return NewScalarTxResult(transactions, len(transactions)), nil
}

func (h *Handler) GetTransactionByHash(request *http.Request) (*Result, *types.Error) {
	hash := chi.URLParam(request, "hash")
	if hash == "" {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "hash is required")
	}
	transaction, err := h.services.GetTransactionByHash(request.Context(), hash)
	if err != nil {
		return nil, err
	}
	return NewResult(transaction), nil
}

func (h *Handler) SearchTransactions(request *http.Request) (*Result, *types.Error) {
	payload, err := parseSearchTransactionsPayload(request)
	if err != nil {
		return nil, err
	}
	transactions, total, err := h.services.SearchTransactions(request.Context(), payload)
	if err != nil {
		return nil, err
	}
	return NewScalarTxResult(transactions, total), nil
}
