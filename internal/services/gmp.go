package services

import (
	"context"
	"net/http"

	"github.com/babylonchain/staking-api-service/internal/db/model"
	"github.com/babylonchain/staking-api-service/internal/types"
)

func (s *Services) GetGMPs(ctx context.Context) ([]*model.GMPDocument, *types.Error) {
	gmps, err := s.DbClient.GetGMPs(ctx)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return gmps, nil
}
