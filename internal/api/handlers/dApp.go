package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/babylonchain/staking-api-service/internal/types"
)

type CreateDAppRequestPayload struct {
	ChainName     string `json:"chain_name"`
	BTCAddressHex string `json:"btc_address_hex"`
	PublicKeyHex  string `json:"public_key_hex"`
}
type UpdateDAppRequestPayload struct {
	ID            string `json:"id"`
	ChainName     string `json:"chain_name"`
	BTCAddressHex string `json:"btc_address_hex"`
	PublicKeyHex  string `json:"public_key_hex"`
}

type IdRequestPayload struct {
	ID string `json:"id"`
}

func parseCreateDAppPayload(request *http.Request) (*CreateDAppRequestPayload, *types.Error) {
	payload := &CreateDAppRequestPayload{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid request payload")
	}
	// // Validate the payload fields - DO it later
	// if !utils.IsValidChainName(payload.ChainName) {
	// 	return nil, types.NewErrorWithMsg(
	// 		http.StatusBadRequest, types.BadRequest, "invalid chain name",
	// 	)
	// }
	// if !utils.IsValidBtcAddress(payload.BTCAddressHex) {
	// 	return nil, types.NewErrorWithMsg(
	// 		http.StatusBadRequest, types.BadRequest, "invalid address hex",
	// 	)
	// }
	// if !utils.IsValidPublickeyHex(payload.PublicKeyHex) {
	// 	return nil, types.NewErrorWithMsg(
	// 		http.StatusBadRequest, types.BadRequest, "invalid public key hex",
	// 	)
	// }
	return payload, nil
}

func parseUpdateDAppPayload(request *http.Request) (*UpdateDAppRequestPayload, *types.Error) {
	payload := &UpdateDAppRequestPayload{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid request payload")
	}
	return payload, nil
}

func parseIdDAppPayload(request *http.Request) (*IdRequestPayload, *types.Error) {
	payload := &IdRequestPayload{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid request payload")
	}
	return payload, nil
}

func (h *Handler) CreateDApp(request *http.Request) (*Result, *types.Error) {
	payload, err := parseCreateDAppPayload(request)
	if err != nil {
		return nil, err
	}

	err = h.services.CreateDApp(request.Context(), payload.ChainName, payload.BTCAddressHex, payload.PublicKeyHex)
	if err != nil {
		return nil, err
	}

	return NewResult(payload), nil
}

func (h *Handler) GetDApp(request *http.Request) (*Result, *types.Error) {
	// FUTURE WORK: Implement pagination
	// paginationKey, err := parsePaginationQuery(request)
	// if err != nil {
	// 	return nil, err
	// }
	dApps, err := h.services.GetDApp(request.Context())
	if err != nil {
		return nil, err
	}
	return NewResult(dApps), nil
}

func (h *Handler) UpdateDApp(request *http.Request) (*Result, *types.Error) {
	payload, err := parseUpdateDAppPayload(request)
	if err != nil {
		return nil, err
	}
	err = h.services.UpdateDApp(request.Context(), payload.ID, payload.ChainName, payload.BTCAddressHex, payload.PublicKeyHex)
	if err != nil {
		return nil, err
	}

	return NewResult(payload), nil
}

func (h *Handler) ToggleDApp(request *http.Request) (*Result, *types.Error) {
	payload, err := parseIdDAppPayload(request)
	if err != nil {
		return nil, err
	}
	err = h.services.ToggleDApp(request.Context(), payload.ID)
	if err != nil {
		return nil, err
	}
	return NewResult(payload), nil
}

func (h *Handler) DeleteDApp(request *http.Request) (*Result, *types.Error) {
	payload, err := parseIdDAppPayload(request)
	if err != nil {
		return nil, err
	}
	err = h.services.DeleteDApp(request.Context(), payload.ID)
	if err != nil {
		return nil, err
	}

	return NewResult("Delete successfully"), nil
}
