package pg

import (
	"context"
	"net/http"

	"github.com/scalarorg/data-models/relayer"
	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (c *PostgresClient) GetTokenSentRelayData(ctx context.Context, options *models.Options) ([]relayer.RelayData, int, *types.Error) {
	var relayDatas []relayer.RelayData
	var totalCount int64

	if options.Size <= 0 {
		options.Size = 10
	}
	if options.Offset < 0 {
		options.Offset = 0
	}

	// Base query with JOIN
	query := c.DB.Model(&relayer.RelayData{}).
		Joins("LEFT JOIN token_sents ON relay_data.id = token_sents.relay_data_id")

	// Apply filters if tx_hash or event_id exists
	if options.TxHash != "" {
		query = query.Where("token_sents.tx_hash = ?", options.TxHash)
	} else if options.EventId != "" {
		query = query.Where("token_sents.event_id = ?", options.EventId)
	}

	// Count total records (only when not searching by specific criteria)
	if options.EventId == "" && options.TxHash == "" {
		if err := query.Count(&totalCount).Error; err != nil {
			return nil, 0, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
		}
	}

	// Execute main query with pagination
	err := query.
		Order("relay_data.created_at DESC").
		Offset(options.Offset).
		Limit(options.Size).
		Find(&relayDatas).Error

	if err != nil {
		return nil, 0, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}

	// For filtered searches, use the length of results as the count
	if options.EventId != "" || options.TxHash != "" {
		totalCount = int64(len(relayDatas))
	}

	return relayDatas, int(totalCount), nil
}
