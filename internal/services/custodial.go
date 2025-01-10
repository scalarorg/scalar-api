package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/pg/models"
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
	err := s.Pg.SaveCustodial(&models.Custodial{
		Name:            params.Name,
		BtcPublicKeyHex: params.BtcPublicKeyHex,
	})
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil
}

func (s *Services) GetCustodials(ctx context.Context) ([]*models.Custodial, *types.Error) {
	custodials, err := s.Pg.GetCustodials()
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return custodials, nil
}

func (s *Services) GetCustodialByName(ctx context.Context, name string) (*models.Custodial, *types.Error) {
	custodial, err := s.Pg.GetCustodialByName(name)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return custodial, nil
}

func (s *Services) CreateCustodialGroup(ctx context.Context, params CreateCustodialGroupServiceParams) *types.Error {
	custodials, err := s.Pg.GetCustodialByNames(params.CustodialNames)
	saveCustodials := make([]models.Custodial, len(custodials))
	for i, custodial := range custodials {
		saveCustodials[i] = *custodial
	}
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	err = s.Pg.SaveCustodialGroup(&models.CustodialGroup{
		Name:       params.Name,
		Quorum:     params.Quorum,
		Custodials: saveCustodials,
	})
	if err != nil {
		return types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return nil
}

func (s *Services) GetCustodialGroups(ctx context.Context) ([]*models.CustodialGroup, *types.Error) {
	custodialGroups, err := s.Pg.GetCustodialGroups()
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return custodialGroups, nil
}

func (s *Services) GetCustodialGroupByName(ctx context.Context, name string) (*models.CustodialGroup, *types.Error) {
	custodialGroup, err := s.Pg.GetCustodialGroupByName(name)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return custodialGroup, nil
}

func (s *Services) GetShortenCustodialGroups(ctx context.Context) ([]*models.ShortenCustodialGroup, *types.Error) {
	shortenCustodialGroups, err := s.Pg.GetShortenCustodialGroups()
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	return shortenCustodialGroups, nil
}
