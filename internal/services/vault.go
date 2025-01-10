package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (s *Services) VaultSearch(ctx context.Context, payload *types.VaultPayload) ([]*models.VaultDocument, *types.Error) {
	vaults, err := s.Pg.Search(ctx, payload)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	if vaults == nil {
		vaults = []*models.VaultDocument{}
	}
	return vaults, nil
}
