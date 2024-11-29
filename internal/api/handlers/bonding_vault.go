package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/types"
)

func parseVaultPayload(request *http.Request) (*types.VaultPayload, *types.Error) {
	payload := &types.VaultPayload{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid gmp request payload")
	}
	return payload, nil
}

// SearchVault godoc
// @Summary Search vaults
// @Description Searches for vaults based on the provided payload
// @Tags vault
// @Accept json
// @Produce json
// @Success 200 {object} PublicResponse[[]vault.VaultDocument] "List of vaults"
// @Router /v1/vault/searchVault [post]
func (h *Handler) SearchVault(request *http.Request) (*Result, *types.Error) {
	vaultPayload, err := parseVaultPayload(request)
	if err != nil {
		return nil, err
	}
	vaults, err := h.services.VaultSearch(request.Context(), vaultPayload)
	if err != nil {
		return nil, err
	}
	return NewResult(vaults), nil
}
