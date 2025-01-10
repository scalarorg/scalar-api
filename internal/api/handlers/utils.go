package handlers

import "net/http"

func NewListResult[T any](data T, total int) *Result {
	res := &ListPublicResponse[T]{Data: data, Total: total}
	return &Result{Data: res, Status: http.StatusOK}
}

type ListPublicResponse[T any] struct {
	Data  T   `json:"data"`
	Total int `json:"total,omitempty"`
}
