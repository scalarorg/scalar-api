package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (s *Services) GMPSearch(ctx context.Context, payload *models.Options) ([]*models.GMPDocument, int, *types.Error) {
	gmps, total, err := s.Pg.GMPSearch(ctx, payload)
	if err != nil {
		return nil, 0, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	if gmps == nil {
		gmps = []*models.GMPDocument{}
	}
	return gmps, total, nil
}
