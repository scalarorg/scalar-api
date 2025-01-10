package pg

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/data-models/relayer"
	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (c *PostgresClient) TransferSearch(ctx context.Context, options *models.Options) ([]*models.TransferDocument, int, *types.Error) {
	relayData, total, err := c.GetTokenSentRelayData(ctx, options)
	if err != nil {
		return nil, 0, err
	}

	fmt.Println("relayData", relayData)

	result, err := c.getTransferByRelayData(ctx, relayData)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (c *PostgresClient) getTransferByRelayData(ctx context.Context, relayData []relayer.RelayData) ([]*models.TransferDocument, *types.Error) {
	log.Info().Msg("getTransferByRelayData")
	log.Info().Msg(fmt.Sprintf("relayData: %+v", relayData))

	messageIds := make([]string, len(relayData))
	for index, event := range relayData {
		messageIds[index] = event.ID
	}
	transfers := make([]*models.TransferDocument, len(relayData))
	for index, relayData := range relayData {
		transfers[index] = &models.TransferDocument{
			ID:               relayData.ID,
			Status:           strconv.Itoa(int(relayData.Status)),
			SimplifiedStatus: string(models.ToReadableStatus(int(relayData.Status))),
		}
	}
	return transfers, nil
}
