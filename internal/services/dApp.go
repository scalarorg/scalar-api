package services

import (
	"context"
	"net/http"

	"github.com/babylonchain/staking-api-service/internal/db/model"
	"github.com/babylonchain/staking-api-service/internal/types"
)

func (s *Services) CreateDApp(ctx context.Context, chainName, btcAddressHex, publicKeyHex string) *types.Error {
	err := s.DbClient.SaveDApp(ctx, chainName, btcAddressHex, publicKeyHex)
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil

}

func (s *Services) GetDApp(ctx context.Context) ([]*model.DAppDocument, *types.Error) {
	dApps, err := s.DbClient.GetDApp(ctx)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return dApps, nil
}

func (s *Services) UpdateDApp(ctx context.Context, ID, chainName, btcAddressHex, publicKeyHex string) *types.Error {
	err := s.DbClient.UpdateDApp(ctx, ID, chainName, btcAddressHex, publicKeyHex)
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil
}

func (s *Services) ToggleDApp(ctx context.Context, ID string) *types.Error {
	err := s.DbClient.ToggleDApp(ctx, ID)
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil
}

func (s *Services) DeleteDApp(ctx context.Context, ID string) *types.Error {
	err := s.DbClient.DeleteDApp(ctx, ID)
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil
}
