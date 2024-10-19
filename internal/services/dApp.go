package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/model"
	"github.com/scalarorg/xchains-api/internal/types"
)

type DAppServiceParams struct {
	ID                   string // TODO: change to ObjectID
	ChainName            string
	BtcAddressHex        string
	PublicKeyHex         string
	SmartContractAddress string
	ChainID              string
	ChainEndpoint        string
	RpcUrl               string
	AccessToken          string
}

func (s *Services) CreateDApp(ctx context.Context, params DAppServiceParams) *types.Error {

	err := s.DbClient.SaveDApp(ctx, &model.DAppDocument{
		ChainName:            params.ChainName,
		BTCAddressHex:        params.BtcAddressHex,
		PublicKeyHex:         params.PublicKeyHex,
		SmartContractAddress: params.SmartContractAddress,
		ChainID:              params.ChainID,
		ChainEndpoint:        params.ChainEndpoint,
		RPCUrl:               params.RpcUrl,
		AccessToken:          params.AccessToken,
	})
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

func (s *Services) UpdateDApp(ctx context.Context, params DAppServiceParams) *types.Error {

	err := s.DbClient.UpdateDApp(ctx, &model.DAppDocument{
		ID:                   params.ID,
		ChainName:            params.ChainName,
		BTCAddressHex:        params.BtcAddressHex,
		PublicKeyHex:         params.PublicKeyHex,
		SmartContractAddress: params.SmartContractAddress,
		ChainID:              params.ChainID,
		ChainEndpoint:        params.ChainEndpoint,
		RPCUrl:               params.RpcUrl,
		AccessToken:          params.AccessToken,
	})

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
