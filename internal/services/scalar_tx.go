package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/types"
)

func (s *Services) GetTransactions(ctx context.Context, payload *types.GetTransactionsRequestPayload) ([]*types.GetTransactionsResponsePayload, *types.Error) {
	height := payload.Height
	transactions, err := s.IndexerAdapter.GetTransactionsByBlockHeight(ctx, height)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	response := make([]*types.GetTransactionsResponsePayload, len(transactions))
	for i, tx := range transactions {
		response[i] = &types.GetTransactionsResponsePayload{
			TxHash: tx.Hash,
		}
	}
	return response, nil
}

func (s *Services) GetTransactionByHash(ctx context.Context, hash string) (*types.GetTransactionByHashResponsePayload, *types.Error) {
	transaction, err := s.IndexerAdapter.GetTransactionByHash(ctx, hash)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	// TODO: Implement missing fields
	return &types.GetTransactionByHashResponsePayload{
		Height:    int(transaction.Block.Height),
		Type:      "",
		Code:      int(transaction.Code),
		Timestamp: transaction.Block.TimeStamp.Unix(),
		GasUsed:   0,
		GasWanted: 0,
	}, nil
}

func (s *Services) SearchTransactions(ctx context.Context, payload *types.SearchTransactionsRequestPayload) ([]*types.GetTransactionsResponsePayload, int, *types.Error) {
	transactions, total, err := s.IndexerAdapter.SearchTransactions(ctx, payload)
	if err != nil {
		return nil, 0, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	response := make([]*types.GetTransactionsResponsePayload, len(transactions))
	for i, tx := range transactions {
		response[i] = &types.GetTransactionsResponsePayload{
			TxHash: tx.Hash,
		}
	}
	return response, total, nil
}
