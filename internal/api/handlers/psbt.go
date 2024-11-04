package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/psbt"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (h *Handler) CreateStakingPsbt(request *http.Request) (*Result, *types.Error) {
	payload := &types.CreateStakingRequestPayload{}
	err := json.NewDecoder(request.Body).Decode(payload)
	if err != nil {
		return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid staking request payload")
	}
	//1. Get AddressUtxos from mempool api base on stakerAddress
	//2. Get feerate from mempool api
	//3. Create PSBT

	unsignedPsbt, err := psbt.CreateUnsignedStakingPsbt(request.Context(), payload)
	return NewResult(unsignedPsbt), nil
}

func (h *Handler) CreateUnstakingPsbt(request *http.Request) (*Result, *types.Error) {
	// payload := &types.CreateUnstakingRequestPayload{}
	// err := json.NewDecoder(request.Body).Decode(payload)
	// if err != nil {
	// 	return nil, types.NewErrorWithMsg(http.StatusBadRequest, types.BadRequest, "invalid staking request payload")
	// }
	// unstakingPsbt, err := psbt.CreateUnstakingPsbt(request.Context(), payload)
	// return NewResult(unstakingPsbt), nil
	return nil, nil
}
