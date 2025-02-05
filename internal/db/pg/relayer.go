package pg

import (
	"context"
	"net/http"

	"github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (c *PostgresClient) ListTokenSents(ctx context.Context, options *models.Options) ([]*chains.TokenSent, int, *types.Error) {
	var tokenSents []*chains.TokenSent
	var totalCount int64

	if options.Size <= 0 {
		options.Size = 10
	}
	if options.Offset < 0 {
		options.Offset = 0
	}

	query := c.DB.Model(&chains.TokenSent{})

	if options.TxHash != "" {
		query = query.Where("tx_hash = ?", options.TxHash)
	} else if options.EventId != "" {
		query = query.Where("event_id = ?", options.EventId)
	}

	if options.EventId == "" && options.TxHash == "" {
		if err := query.Count(&totalCount).Error; err != nil {
			return nil, 0, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
		}
	}

	err := query.
		Order("created_at DESC").
		Offset(options.Offset).
		Limit(options.Size).
		Find(&tokenSents).Error

	if err != nil {
		return nil, 0, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}

	// For filtered searches, use the length of results as the count
	if options.EventId != "" || options.TxHash != "" {
		totalCount = int64(len(tokenSents))
	}

	return tokenSents, int(totalCount), nil
}

func (c *PostgresClient) ListEvmToBTCTransfers(ctx context.Context, options *models.Options) ([]*chains.ContractCallWithToken, int, *types.Error) {
	var contractCallWithTokens []*chains.ContractCallWithToken
	var totalCount int64

	if options.Size <= 0 {
		options.Size = 10
	}
	if options.Offset < 0 {
		options.Offset = 0
	}

	query := c.DB.Model(&chains.ContractCallWithToken{})

	if options.TxHash != "" {
		query = query.Where("tx_hash = ?", options.TxHash)
	} else if options.EventId != "" {
		query = query.Where("event_id = ?", options.EventId)
	}

	// use LIKE to match evm and bitcoin
	query = query.Where("source_chain LIKE ?", "%evm%")
	query = query.Where("destination_chain LIKE ?", "%bitcoin%")

	if options.EventId == "" && options.TxHash == "" {
		if err := query.Count(&totalCount).Error; err != nil {
			return nil, 0, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
		}
	}

	err := query.
		Order("created_at DESC").
		Offset(options.Offset).
		Limit(options.Size).
		Find(&contractCallWithTokens).Error

	if err != nil {
		return nil, 0, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}

	// For filtered searches, use the length of results as the count
	if options.EventId != "" || options.TxHash != "" {
		totalCount = int64(len(contractCallWithTokens))
	}

	return contractCallWithTokens, int(totalCount), nil
}
