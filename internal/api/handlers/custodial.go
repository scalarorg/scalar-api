package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/postgres/models"
	"github.com/scalarorg/xchains-api/internal/services"
	"github.com/scalarorg/xchains-api/internal/types"
)

func parseCreateCustodialPayload(request *http.Request) (*services.CreateCustodialServiceParams, *types.Error) {
	payload := &services.CreateCustodialServiceParams{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid request payload")
	}
	return payload, nil
}

func parseCreateCustodialGroupPayload(request *http.Request) (*services.CreateCustodialGroupServiceParams, *types.Error) {
	payload := &services.CreateCustodialGroupServiceParams{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid request payload")
	}
	return payload, nil
}

func (h *Handler) CreateCustodial(request *http.Request) (*Result, *types.Error) {
	payload, err := parseCreateCustodialPayload(request)
	if err != nil {
		return nil, err
	}

	err = h.services.CreateCustodial(request.Context(), *payload)
	if err != nil {
		return nil, err
	}

	return NewResult(payload), nil
}

func (h *Handler) GetCustodial(request *http.Request) (*Result, *types.Error) {
	custodials, err := h.services.GetCustodial(request.Context())
	if err != nil {
		return nil, err
	}
	if custodials == nil {
		custodials = []*models.Custodial{}
	}
	return NewResult(custodials), nil
}

func (h *Handler) GetCustodialByName(request *http.Request) (*Result, *types.Error) {
	custodial, err := h.services.GetCustodialByName(request.Context(), request.URL.Query().Get("name"))
	if err != nil {
		return nil, err
	}
	if custodial == nil {
		custodial = &models.Custodial{}
	}
	return NewResult(custodial), nil
}

func (h *Handler) CreateCustodialGroup(request *http.Request) (*Result, *types.Error) {
	payload, err := parseCreateCustodialGroupPayload(request)
	if err != nil {
		return nil, err
	}

	err = h.services.CreateCustodialGroup(request.Context(), *payload)
	if err != nil {
		return nil, err
	}

	return NewResult(payload), nil
}

func (h *Handler) GetCustodialGroupByName(request *http.Request) (*Result, *types.Error) {
	custodialGroup, err := h.services.GetCustodialGroupByName(request.Context(), request.URL.Query().Get("name"))
	if err != nil {
		return nil, err
	}
	return NewResult(custodialGroup), nil
}
