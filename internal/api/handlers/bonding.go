// internal/api/handlers/bonding.go

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/types"
	"github.com/scalarorg/xchains-api/internal/utils"
)

// BondingRequest represents the expected request body for creating a bonding transaction.
type BondingRequestPayload struct {
	Amount                   int    `json:"amount"`
	BondHolderAddress        string `json:"bond_holder_address"`
	BondHolderPublicKey      string `json:"bond_holder_public_key"`
	ProtocolPublicKey        string `json:"protocol_public_key"`
	DestChainId              string `json:"dest_chain_id"`
	DestUserAddress          string `json:"dest_user_address"`
	DestSmartContractAddress string `json:"dest_smart_contract_address"`
}

func parseBondTransactionRequestPayload(request *http.Request) (*BondingRequestPayload, *types.Error) {
	payload := &BondingRequestPayload{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid request payload")
	}
	// Validate the payload fields
	if !utils.IsValidAddress(payload.BondHolderAddress) {
		return nil, types.NewErrorWithMsg(
			http.StatusBadRequest, types.BadRequest, "invalid bond holder address",
		)
	}
	if !utils.IsValidPublicKey(payload.BondHolderPublicKey) {
		return nil, types.NewErrorWithMsg(
			http.StatusBadRequest, types.BadRequest, "invalid bond holder public key",
		)
	}
	if !utils.IsValidPublicKey(payload.ProtocolPublicKey) {
		return nil, types.NewErrorWithMsg(
			http.StatusBadRequest, types.BadRequest, "invalid protocol public key",
		)
	}

	return payload, nil
}

// CreateBondingTransaction handles the HTTP request for creating a bonding transaction.
func (h *Handler) BondTransaction(request *http.Request) (*Result, *types.Error) {
	payload, err := parseBondTransactionRequestPayload(request)
	if err != nil {
		return nil, err
	}
	psbt, bondErr := h.services.BondTransaction(payload.Amount, payload.BondHolderAddress, payload.BondHolderPublicKey, payload.ProtocolPublicKey, payload.DestChainId, payload.DestUserAddress, payload.DestSmartContractAddress)
	if bondErr != nil {
		return nil, bondErr
	}
	return NewResult(psbt), nil
}
