package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/postgres/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

type CreateCustodialServiceParams struct {
	Name            string `json:"name"`
	BtcPublicKeyHex string `json:"btc_public_key_hex"`
}

type CreateCustodialGroupServiceParams struct {
	Name           string   `json:"name"`
	Quorum         uint     `json:"quorum"`
	CustodialNames []string `json:"custodial_names"`
}

func (s *Services) CreateCustodial(ctx context.Context, params CreateCustodialServiceParams) *types.Error {
	err := s.ScalarClient.SaveCustodial(&models.Custodial{
		Name:            params.Name,
		BtcPublicKeyHex: params.BtcPublicKeyHex,
	})
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil
}

func (s *Services) GetCustodial(ctx context.Context) ([]*models.Custodial, *types.Error) {
	custodials, err := s.ScalarClient.GetCustodials()
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return custodials, nil
}

func (s *Services) GetCustodialByName(ctx context.Context, name string) (*models.Custodial, *types.Error) {
	custodial, err := s.ScalarClient.GetCustodialByName(name)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return custodial, nil
}

func (s *Services) CreateCustodialGroup(ctx context.Context, params CreateCustodialGroupServiceParams) *types.Error {
	custodials, err := s.ScalarClient.GetCustodialByNames(params.CustodialNames)
	saveCustodials := make([]models.Custodial, len(custodials))
	for i, custodial := range custodials {
		saveCustodials[i] = *custodial
	}
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	err = s.ScalarClient.SaveCustodialGroup(&models.CustodialGroup{
		Name:       params.Name,
		Quorum:     params.Quorum,
		Custodials: saveCustodials,
	})
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil
}

func (s *Services) GetCustodialGroupByName(ctx context.Context, name string) (*models.CustodialGroup, *types.Error) {
	custodialGroup, err := s.ScalarClient.GetCustodialGroupByName(name)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return custodialGroup, nil
}
