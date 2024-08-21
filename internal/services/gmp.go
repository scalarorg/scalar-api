package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/model"
	"github.com/scalarorg/xchains-api/internal/db/postgres"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (s *Services) GMPSearch(ctx context.Context, payload *types.GmpPayload) ([]*postgres.GMPDocument, *types.Error) {
	gmps, err := s.PgClient.GMPSearch(ctx, payload)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	if gmps == nil {
		gmps = []*postgres.GMPDocument{}
	}
	return gmps, nil
}

func (s *Services) GmpGetContracts(ctx context.Context) ([]*model.GMPDocument, *types.Error) {
	gmps, err := s.DbClient.GetGMPs(ctx)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return gmps, nil
}

func (s *Services) GmpGetConfigurations(ctx context.Context) ([]*model.GMPDocument, *types.Error) {
	gmps, err := s.DbClient.GetGMPs(ctx)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return gmps, nil
}
