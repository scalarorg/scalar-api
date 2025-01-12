package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (s *Services) TokenSearchTransfers(ctx context.Context, options *models.Options) ([]*models.TransferDocument, int, *types.Error) {
	transfers, total, err := s.Pg.TokenSearchTransfers(ctx, options)
	if err != nil {
		return nil, 0, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	if transfers == nil {
		transfers = []*models.TransferDocument{}
	}
	return transfers, total, nil
}
