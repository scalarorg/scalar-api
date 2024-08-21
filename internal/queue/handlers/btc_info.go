package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	queueClient "github.com/scalarorg/staking-queue-client/client"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (h *QueueHandler) BtcInfoHandler(ctx context.Context, messageBody string) *types.Error {
	var btcInfo queueClient.BtcInfoEvent
	err := json.Unmarshal([]byte(messageBody), &btcInfo)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Failed to unmarshal the message body into btcInfo")
		return types.NewError(http.StatusBadRequest, types.BadRequest, err)
	}

	statsErr := h.Services.ProcessBtcInfoStats(
		ctx, btcInfo.Height, btcInfo.ConfirmedTvl, btcInfo.UnconfirmedTvl,
	)
	if statsErr != nil {
		log.Error().Err(statsErr).Msg("Failed to process unconfirmed tvl stats")
		return types.NewInternalServiceError(statsErr)
	}
	return nil
}
