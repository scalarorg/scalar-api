package psbt

import (
	"context"

	"github.com/scalarorg/xchains-api/internal/types"
)

func CreateUnsignedStakingPsbt(ctx context.Context, stakingRequest *types.CreateStakingRequestPayload) (*types.StakingPsbt, error) {
	// // create the p2tr script pubkey from public key
	// //const p2trScriptPubKey = publicKeyToP2trScript(stakerPubkey, network);
	// // 1. Create the staking output
	// tag := "staking"
	// version := 0
	// stakerPubkey, err := hex.DecodeString(stakingRequest.StakerPubkey)
	// if err != nil {
	// 	return nil, err
	// }
	// protocolPubkey, err := hex.DecodeString(stakingRequest.ServicePubkey)
	// if err != nil {
	// 	return nil, err
	// }
	// custodialPubkeys := []byte{}
	// // for _, custodialPubkey := range stakingRequest.CustodialPubkeys {
	// // 	custodialPubkeyBytes, err := hex.DecodeString(custodialPubkey)
	// // 	if err != nil {
	// // 		return nil, err
	// // 	}
	// // 	custodialPubkeys = append(custodialPubkeys, custodialPubkeyBytes)
	// // }
	// custodialQuorum := 3
	// haveOnlyCustodial := false
	// dstChainId, err := strconv.ParseUint(stakingRequest.DstChainId, 10, 64)
	// dstSmartContractAddress, err := hex.DecodeString(stakingRequest.DstSmartContractAddress)
	// dstUserAddress, err := hex.DecodeString(stakingRequest.DstUserAddress)
	// outputs, err := psbtFfi.CreateStakingPsbt(
	// 	tag,
	// 	version,
	// 	stakingRequest.StakingAmount,
	// 	stakerPubkey,
	// 	protocolPubkey,
	// 	custodialPubkeys,
	// 	custodialQuorum,
	// 	haveOnlyCustodial,
	// 	dstChainId,
	// 	dstSmartContractAddress,
	// 	dstUserAddress,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	// return outputs, nil
	return nil, nil
}

// func BuildUnsignedStakingPsbt(
//   tag: string,
//   version: number,
//   network: Network,
//   stakerAddress: string,
//   stakerPubkey: Uint8Array,
//   protocolPubkey: Uint8Array,
//   custodialPubkeys: Uint8Array,
//   custodialQuorum: number,
//   haveOnlyCustodial: boolean,
//   dstChainId: bigint,
//   dstSmartContractAddress: Uint8Array,
//   dstUserAddress: Uint8Array,
//   addressUtxos: AddressTxsUtxo[],
//   feeRate: number,
//   stakingAmount: bigint,
//   rbf: boolean = true
// )  {
//   // create the p2tr script pubkey from public key
//   //const p2trScriptPubKey = publicKeyToP2trScript(stakerPubkey, network);
//   // 1. Create the staking output
//   const outputs = buildStakingOutput(
//     tag,
//     version,
//     stakingAmount,
//     stakerPubkey,
//     protocolPubkey,
//     custodialPubkeys,
//     custodialQuorum,
//     haveOnlyCustodial,
//     dstChainId,
//     dstSmartContractAddress,
//     dstUserAddress
//   );

//   // 2. Get the selected utxos and fees
//   const inputByAddress = prepareExtraInputByAddress(
//     stakerAddress,
//     stakerPubkey,
//     network
//   );

//   const regularUTXOs: UTXO[] = addressUtxos.map(
//     ({ txid, vout, value }: AddressTxsUtxo) => ({
//       txid,
//       vout,
//       value,
//     })
//   );
//   const { selectedUTXOs, fee } = getStakingTxInputUTXOsAndFees(
//     network,
//     regularUTXOs,
//     inputByAddress.outputScriptSize,
//     Number(stakingAmount),
//     feeRate,
//     outputs
//   );
//   // 3. Create the psbt
//   const { psbt, fee: estimatedFee } = createStakingPsbt(
//     network,
//     inputByAddress,
//     selectedUTXOs,
//     outputs,
//     Number(stakingAmount),
//     fee,
//     stakerAddress,
//     rbf
//   );
//   return {
//     psbt,
//     fee: estimatedFee,
//   };
// };
