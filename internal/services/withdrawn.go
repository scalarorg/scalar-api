package services

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/xchains-api/internal/db"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (s *Services) TransitionToWithdrawnState(
	ctx context.Context, stakingTxHashHex string,
) *types.Error {
	err := s.DbClient.TransitionToWithdrawnState(ctx, stakingTxHashHex)
	if err != nil {
		if ok := db.IsNotFoundError(err); ok {
			log.Ctx(ctx).Warn().Str("stakingTxHashHex", stakingTxHashHex).Err(err).Msg("delegation not found or no longer eligible for withdraw")
			return types.NewErrorWithMsg(http.StatusForbidden, types.NotFound, "delegation not found or no longer eligible for withdraw")
		}
		log.Ctx(ctx).Error().Str("stakingTxHashHex", stakingTxHashHex).Err(err).Msg("failed to transition to withdrawn state")
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil
}
