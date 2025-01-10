package pg

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (c *PostgresClient) FindEventsByType(ctx context.Context, options *models.Options) ([]models.BlockEvent, *types.Error) {
	var events []models.BlockEvent
	// this can potentially be optimized by getting max first and selecting it (this gets translated into a select * limit 1)
	// c.Client.Joins("inner join txes on tx_messages.tx_id = txes.id").Find(&txMsgs)
	if options.Size <= 0 {
		options.Size = 10
	}
	if options.Offset < 0 {
		options.Offset = 0
	}
	result := c.DB.Limit(options.Size).Offset(options.Offset).InnerJoins("BlockEventType", c.DB.Where(&models.BlockEventType{Type: options.EventType})).Find(&events)
	if result.Error != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, result.Error)
	}

	return events, nil
}
