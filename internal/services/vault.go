package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/vault"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (s *Services) VaultSearch(ctx context.Context, payload *types.VaultPayload) ([]*vault.VaultDocument, *types.Error) {
	vaults, err := s.VaultClient.Search(ctx, payload)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	if vaults == nil {
		vaults = []*vault.VaultDocument{}
	}
	return vaults, nil
}
