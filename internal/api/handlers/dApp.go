package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/services"
	"github.com/scalarorg/xchains-api/internal/types"
)

type CreateDAppRequestPayload struct {
	ChainName            string `json:"chain_name"`
	BTCAddressHex        string `json:"btc_address_hex"`
	PublicKeyHex         string `json:"public_key_hex"`
	SmartContractAddress string `json:"smart_contract_address"` //UPDATE: New field
	ChainID              string `json:"chain_id"`
	ChainEndpoint        string `json:"chain_endpoint"`
	RPCUrl               string `json:"rpc_url"`
	AccessToken          string `json:"access_token"`
	TokenContractAddress string `json:"token_contract_address"`
	CustodialGroupID     uint   `json:"custodial_group_id"`
}
type UpdateDAppRequestPayload struct {
	ID                   string `json:"id"`
	ChainName            string `json:"chain_name"`
	BTCAddressHex        string `json:"btc_address_hex"`
	PublicKeyHex         string `json:"public_key_hex"`
	SmartContractAddress string `json:"smart_contract_address"` //UPDATE: New field
	ChainID              string `json:"chain_id"`
	ChainEndpoint        string `json:"chain_endpoint"`
	RPCUrl               string `json:"rpc_url"`
	AccessToken          string `json:"access_token"`
	TokenContractAddress string `json:"token_contract_address"`
	CustodialGroupID     uint   `json:"custodial_group_id"`
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

// CreateDApp godoc
// @Summary Create dApp
// @Description Creates a new dApp
// @Tags dApp
// @Accept json
// @Produce json
// @Success 200 {object} PublicResponse[CreateDAppRequestPayload] "Created dApp"
// @Router /v1/dApp [post]
func (h *Handler) CreateDApp(request *http.Request) (*Result, *types.Error) {
	payload, err := parseCreateDAppPayload(request)
	if err != nil {
		return nil, err
	}

	params := services.DAppServiceParams{
		ChainName:            payload.ChainName,
		BtcAddressHex:        payload.BTCAddressHex,
		PublicKeyHex:         payload.PublicKeyHex,
		SmartContractAddress: payload.SmartContractAddress,
		ChainID:              payload.ChainID,
		ChainEndpoint:        payload.ChainEndpoint,
		RpcUrl:               payload.RPCUrl,
		AccessToken:          payload.AccessToken,
		TokenContractAddress: payload.TokenContractAddress,
		CustodialGroupID:     payload.CustodialGroupID,
	}

	err = h.services.CreateDApp(request.Context(), params)

	if err != nil {
		return nil, err
	}

	return NewResult(payload), nil
}

// GetDApp godoc
// @Summary Get dApps
// @Description Retrieves all dApps
// @Tags dApp
// @Produce json
// @Success 200 {object} PublicResponse[[]models.DApp] "List of dApps"
// @Router /v1/dApp [get]
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
	if dApps == nil {
		dApps = []*models.DApp{}
	}
	return NewResult(dApps), nil
}

// UpdateDApp godoc
// @Summary Update dApp
// @Description Updates a dApp
// @Tags dApp
// @Produce json
// @Success 200 {object} PublicResponse[UpdateDAppRequestPayload] "Updated dApp"
// @Router /v1/dApp [put]
func (h *Handler) UpdateDApp(request *http.Request) (*Result, *types.Error) {
	payload, err := parseUpdateDAppPayload(request)
	if err != nil {
		return nil, err
	}

	params := services.DAppServiceParams{
		ID:                   payload.ID,
		ChainName:            payload.ChainName,
		BtcAddressHex:        payload.BTCAddressHex,
		PublicKeyHex:         payload.PublicKeyHex,
		SmartContractAddress: payload.SmartContractAddress,
		ChainID:              payload.ChainID,
		ChainEndpoint:        payload.ChainEndpoint,
		RpcUrl:               payload.RPCUrl,
		AccessToken:          payload.AccessToken,
		TokenContractAddress: payload.TokenContractAddress,
		CustodialGroupID:     payload.CustodialGroupID,
	}

	err = h.services.UpdateDApp(request.Context(), params)
	if err != nil {
		return nil, err
	}

	return NewResult(payload), nil
}

// ToggleDApp godoc
// @Summary Toggle dApp
// @Description Toggles a dApp
// @Tags dApp
// @Produce json
// @Success 200 {object} PublicResponse[IdRequestPayload] "Toggled dApp"
// @Router /v1/dApp [patch]
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

// DeleteDApp godoc
// @Summary Delete dApp
// @Description Deletes a dApp
// @Tags dApp
// @Produce json
// @Success 200 {string} string "Delete successfully"
// @Router /v1/dApp [delete]
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
