package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/pg/models"
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
	TokenContractAddress string
	CustodialGroupID     uint
}

func (s *Services) CreateDApp(ctx context.Context, params DAppServiceParams) *types.Error {
	err := s.Pg.SaveDApp(&models.DApp{
		ChainName:            params.ChainName,
		BTCAddressHex:        params.BtcAddressHex,
		PublicKeyHex:         params.PublicKeyHex,
		SmartContractAddress: params.SmartContractAddress,
		ChainID:              params.ChainID,
		ChainEndpoint:        params.ChainEndpoint,
		RPCUrl:               params.RpcUrl,
		AccessToken:          params.AccessToken,
		TokenContractAddress: params.TokenContractAddress,
		CustodialGroupID:     params.CustodialGroupID,
	})
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil

}

func (s *Services) GetDApp(ctx context.Context) ([]*models.DApp, *types.Error) {
	dApps, err := s.Pg.GetDApps()
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return dApps, nil
}

func (s *Services) UpdateDApp(ctx context.Context, params DAppServiceParams) *types.Error {

	err := s.Pg.UpdateDApp(&models.DApp{
		ID:                   params.ID,
		ChainName:            params.ChainName,
		BTCAddressHex:        params.BtcAddressHex,
		PublicKeyHex:         params.PublicKeyHex,
		SmartContractAddress: params.SmartContractAddress,
		ChainID:              params.ChainID,
		ChainEndpoint:        params.ChainEndpoint,
		RPCUrl:               params.RpcUrl,
		AccessToken:          params.AccessToken,
		TokenContractAddress: params.TokenContractAddress,
		CustodialGroupID:     params.CustodialGroupID,
	})

	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil
}

func (s *Services) ToggleDApp(ctx context.Context, ID string) *types.Error {
	err := s.Pg.ToggleDApp(ID)
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil
}

func (s *Services) DeleteDApp(ctx context.Context, ID string) *types.Error {
	err := s.Pg.DeleteDApp(ID)
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil
}
