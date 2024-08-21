package postgres

import (
	"context"
	"fmt"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/types"
)

func (c *PostgresClient) GMPSearch(ctx context.Context, payload *types.GmpPayload) ([]*GMPDocument, *types.Error) {
	var txMsgs []TxMessage
	// this can potentially be optimized by getting max first and selecting it (this gets translated into a select * limit 1)
	// c.Client.Joins("inner join txes on tx_messages.tx_id = txes.id").Find(&txMsgs)
	if payload.Size <= 0 {
		payload.Size = 10
	}
	result := c.Client.Limit(payload.Size).Preload("Tx").Find(&txMsgs)
	if result.Error != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, result.Error)
	}
	gmps := make([]*GMPDocument, len(txMsgs))
	for index, msg := range txMsgs {
		gmps[index] = &GMPDocument{
			Call: GMPStepDocument{
				ID:              msg.Tx.Hash,
				TransactionHash: msg.Tx.Hash,
			},
		}
		fmt.Printf("gmp %v create from %v", gmps[index].Call.ID, msg.Tx)
	}
	return gmps, nil
}
