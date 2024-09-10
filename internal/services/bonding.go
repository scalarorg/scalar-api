package services

import (
	"github.com/scalarorg/xchains-api/internal/types"
)

type psbt struct {
	PSBTHex string `json:"psbt"`
}

func (s *Services) BondTransaction(amount int, bondHolderAddress string, bondHolderPublicKey string, protocolPublicKey string, destChainId string, destUserAddress string, destSmartContractAddress string) (types.Staker, *types.Error) {
	staker := types.Staker{
		StakerAddress:             bondHolderAddress,
		StakerPubkey:              bondHolderPublicKey,
		ProtocolPubkey:            protocolPublicKey,
		CovenantPubkeys:           types.CovenantParams.CovenantPubkeys,
		Qorum:                     types.CovenantParams.Qorum,
		Tag:                       types.CovenantParams.Tag,
		Version:                   types.CovenantParams.Version,
		ChainID:                   destChainId,
		ChainIdUserAddress:        destUserAddress,
		ChainSmartContractAddress: destSmartContractAddress,
		MintingAmount:             amount,
	}
	// psbt := psbt{PSBTHex: "abcabc123123"}

	return staker, nil
}
