package pg

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/scalarorg/data-models/relayer"
	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

const (
	EVENT_TYPE_MESSAGE                  string = "message"
	EVENT_TYPE_CONTRACT_CALL            string = "CallContract"
	EVENT_TYPE_CONTRACT_CALL_WITH_TOKEN string = "CallContractWithToken"
	EVENT_TYPE_CONTRACT_CALL_APPROVED   string = "scalar.chains.v1beta1.CallContractApproved"
	EVENT_TYPE_MESSAGE_PROCESSING       string = "scalar.chains.v1beta1.MessageProcessing"
	EVENT_TYPE_MESSAGE_EXECUTED         string = "scalar.chains.v1beta1.MessageExecuted"
	EVENT_TYPE_XCHAINS_CONFIRM          string = "xchains.confirm"
)

func normalizeEventType(eventType string) string {
	parts := strings.Split(eventType, ".")
	return parts[len(parts)-1]
}

func (c *PostgresClient) GMPSearch(ctx context.Context, options *models.Options) ([]*models.GMPDocument, int, *types.Error) {
	relayData, total, err := c.GetRelayerData(ctx, options)
	if err != nil {
		return nil, 0, err
	}

	result, err := c.getGMPByRelayData(ctx, relayData)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (c *PostgresClient) getGMPByRelayData(ctx context.Context, relayDatas []relayer.RelayData) ([]*models.GMPDocument, *types.Error) {
	messageIds := make([]string, len(relayDatas))
	for index, event := range relayDatas {
		messageIds[index] = event.ID
	}
	// options := &postgres.Options{}
	// blockEvents, err := c.indexer.FindBlockEventsByAttribute(ctx, "event_id", messageIds, options)
	// if err != nil {
	// 	return nil, err
	// }
	// var eventIds []uint
	// for _, event := range blockEvents {
	// 	eventIds = append(eventIds, event.ID)
	// }
	// eventAttributes, err := c.indexer.FindEventAttributes(ctx, eventIds)
	// fmt.Printf("EventAttributes %v", eventAttributes)
	mapGmps := make(map[string]*models.GMPDocument)
	mapRelayDatas := make(map[string]relayer.RelayData)
	gmps := make([]*models.GMPDocument, len(relayDatas))
	for index, relayData := range relayDatas {
		gmps[index] = &models.GMPDocument{
			ID: relayData.ID,
			// From:   relayData.From,
			// To:     relayData.To,
			// Status: relayData.Status.String,
			// CreatedAt: relayData.CreatedAt,
			// UpdatedAt: relayData.UpdatedAt,
			Status:           strconv.Itoa(int(relayData.Status)),
			SimplifiedStatus: string(models.ToReadableStatus(int(relayData.Status))),
		}
		createCallContract(gmps[index], &relayData)
		createCallContractApproved(gmps[index], &relayData)
		mapRelayDatas[relayData.ID] = relayData
		mapGmps[relayData.ID] = gmps[index]
	}
	// for _, event := range blockEvents {
	// 	if attributes, ok := eventAttributes[event.ID]; ok {
	// 		if event_id, ok := attributes["event_id"]; ok {
	// 			eventId := parseAttributeValue(event_id).(string)
	// 			relayData, rok := mapRelayDatas[eventId]
	// 			gmp, gok := mapGmps[eventId]
	// 			if rok && gok {
	// 				createGMPDocument(gmp, &relayData, &event, attributes)
	// 			}
	// 		}
	// 	}
	// }

	// debugMsg := "0x04936b1d1304e2d2b5cb841126b13507eeb36a3482fff9a6f16571062e00a3cd-0"
	// messageIds = []string{debugMsg}
	// mapCallContracts, err := c.relayer.GetCallContractParams(ctx, messageIds)
	// if err == nil {
	// 	CallContract := mapCallContracts[debugMsg]
	// 	fmt.Printf("Contract call %v\n", CallContract)
	// 	// for _, gmp := range gmps {
	// 	// 	// gmp.Call.TransactionHash = CallContract.PayloadHash
	// 	// 	// if CallContract, ok := mapCallContracts[gmp.Call.ID]; ok {
	// 	// 	// 	fmt.Printf("Contract call %v", CallContract)
	// 	// 	// 	// gmp.Call.Params = CallContract.Params
	// 	// 	// }

	// 	// }
	// } else {
	// 	fmt.Printf("Get contract call with Error %v", err)
	// }
	return gmps, nil
}

func createCallContract(gmp *models.GMPDocument, relayData *relayer.RelayData) {
	from := "relayData.CallContract.Sender.String"
	// Nov 11, Taivv: SenderAddress is stored using electrum indexer, get from first txin.script_pubkey
	// if relayData.CallContract.StakerPublicKey.String != "" {
	// 	stakerAddress, err := config.GetTaprootAddress(relayData.From.String, relayData.CallContract.StakerPublicKey.String)
	// 	if err != nil {
	// 		stakerAddress = relayData.From.String
	// 		fmt.Printf("[ERROR] Failed to create Taproot address: %v", err)
	// 	}
	// 	from = stakerAddress
	// } else {
	// 	from = relayData.CallContract.SenderAddress.String
	// }

	call := models.GMPStepDocument{
		ID:              relayData.ID,
		Chain:           relayData.From,
		Event:           normalizeEventType(EVENT_TYPE_CONTRACT_CALL),
		ContractAddress: relayData.CallContract.DestContractAddress,
		Transaction: models.TransactionDocument{
			Hash: relayData.CallContract.TxHash,
			From: from,
		},
		ReturnValues: models.ReturnValuesDocument{
			DestinationChain: relayData.To,
		},
		BlockNumber:     uint64(relayData.CallContract.BlockNumber),
		TransactionHash: relayData.CallContract.TxHash,
		//XChains Todo: Created
		BlockTimestamp: relayData.CreatedAt.Unix(),
		LogIndex:       uint(relayData.CallContract.LogIndex),
	}

	if relayData.CallContract.DestContractAddress != "" {
		call.ContractAddress = relayData.CallContract.DestContractAddress
		call.ReturnValues.Sender = relayData.CallContract.SourceAddress
		call.ReturnValues.PayloadHash = relayData.CallContract.PayloadHash
		call.ReturnValues.Payload = hex.EncodeToString(relayData.CallContract.Payload)
		call.ReturnValues.SourceAddress = relayData.CallContract.SourceAddress
		call.ReturnValues.DestinationContractAddress = relayData.CallContract.DestContractAddress
		call.ReturnValues.DestinationChain = relayData.To
		call.ReturnValues.ContractAddress = relayData.CallContract.DestContractAddress
	}
	// if relayData.CallContractWithToken.ContractAddress.Valid {
	// 	call.ContractAddress = relayData.CallContractWithToken.ContractAddress.String
	// 	call.ReturnValues.Sender = relayData.CallContractWithToken.SourceAddress.String
	// 	call.ReturnValues.PayloadHash = relayData.CallContractWithToken.PayloadHash.String
	// 	call.ReturnValues.Payload = relayData.CallContractWithToken.Payload.String
	// 	call.ReturnValues.SourceAddress = relayData.CallContractWithToken.SourceAddress.String
	// 	call.ReturnValues.DestinationContractAddress = relayData.CallContractWithToken.ContractAddress.String
	// 	call.ReturnValues.DestinationChain = relayData.To.String
	// 	call.ReturnValues.ContractAddress = relayData.CallContractWithToken.ContractAddress.String
	// }
	gmp.Call = call
	// fmt.Printf("CallContract %v\n", relayData.CallContract)
	// fmt.Printf("CallContractWithToken %v\n", relayData.CallContractWithToken)
	// fmt.Printf("Call chain %s\n", call.Chain)
}

func createCallContractApproved(gmp *models.GMPDocument, relayData *relayer.RelayData) {
	approved := models.GMPStepDocument{
		//Todo: Fetch blockhash/Number/Timestamp?
		BlockHash: "",
		// BlockTimestamp: 0,
		Event: normalizeEventType(EVENT_TYPE_CONTRACT_CALL_APPROVED),
		//Todo: Fill ChainType by ChainID/ChainName
		ChainType:       "",
		Address:         relayData.CallContract.SourceAddress,
		ContractAddress: relayData.CallContract.DestContractAddress,
	}

	// if relayData.CallContract..SourceTxHash.Valid {
	// 	approved.ID = relayData.CallContract.CallContractApproved.ID
	// 	approved.BlockNumber = uint64(relayData.CallContract.CallContractApproved.BlockNumber.Int32)
	// 	// approved.BlockTimestamp = relayData.CallContract.CallContractApproved.CreatedAt.Time.Unix()
	// 	//approved.BlockHash = relayData.CallContract.CallContractApproved.BlockHash.String
	// 	approved.Chain = relayData.CallContract.CallContractApproved.SourceChain.String
	// 	//approved.Address = relayData.CallContract.CallContractApproved.SourceAddress.String
	// 	//approved.ContractAddress = relayData.CallContract.CallContractApproved.ContractAddress.String
	// 	approved.TransactionHash = relayData.CallContract.CallContractApproved.TxHash.String
	// 	approved.ReturnValues.SourceChain = relayData.CallContract.CallContractApproved.SourceChain.String
	// 	approved.ReturnValues.SourceEventIndex = string(relayData.CallContract.CallContractApproved.SourceEventIndex.Int64)
	// 	approved.ReturnValues.SourceTxHash = relayData.CallContract.CallContractApproved.SourceTxHash.String
	// 	approved.ReturnValues.SourceAddress = relayData.CallContract.CallContractApproved.SourceAddress.String
	// 	approved.ReturnValues.ContractAddress = relayData.CallContract.CallContractApproved.ContractAddress.String
	// 	approved.ReturnValues.PayloadHash = relayData.CallContract.CallContractApproved.PayloadHash.String
	// 	approved.ReturnValues.CommandID = relayData.CallContract.CallContractApproved.CommandId.String

	// }
	gmp.Approved = approved
}

func createGMPDocument(gmp *models.GMPDocument, relayData *relayer.RelayData, event *models.BlockEvent, attribute models.MapBlockEventAttributes) {
	gmp.ID = parseAttributeValue(attribute["event_id"]).(string)
	gmp.CommandID = parseAttributeValue(attribute["command_id"]).(string)
	// fmt.Printf("Event id %s\n", attribute["event_id"])
	switch event.BlockEventType.Type {
	case EVENT_TYPE_CONTRACT_CALL_APPROVED:
		gmp.Approved = createApprovedEvent(relayData, event, attribute)
	}
	// gmp.Call = createCallContract(event, attribute)
	// gmp.Confirm = createConfirmEvent(event, attribute)
	// gmp.Approved = createApprovedEvent(event, attribute)
	// gmp.Executed = createExecuted(event, attribute)
}

func createConfirmEvent(event *models.BlockEvent, attribute models.MapBlockEventAttributes) models.ConfirmDocument {
	// eventId := parseAttributeValue(attribute["event_id"]).(string)
	confirm := models.ConfirmDocument{
		SourceChain:           "",
		PollId:                "",
		BlockNumber:           0,
		BlockTimestamp:        0,
		TransactionIndex:      0,
		SourceTransactionHash: "",
		Event:                 normalizeEventType(EVENT_TYPE_XCHAINS_CONFIRM),
		TransactionHash:       "",
	}
	return confirm
}

// Load data from event CallContractApproved
func createApprovedEvent(relayData *relayer.RelayData, event *models.BlockEvent, attribute models.MapBlockEventAttributes) models.GMPStepDocument {
	eventId := parseAttributeValue(attribute["event_id"]).(string)
	//Todo fill address and contract address
	address := ""
	contractAddress := "" // The same as address
	approved := models.GMPStepDocument{
		//Todo: TxHash_TransactionIndex_logIndex
		ID: eventId,
		//Todo: Fetch blockhash/Number/Timestamp?
		BlockHash:      "",
		BlockNumber:    0,
		BlockTimestamp: 0,
		Event:          normalizeEventType(EVENT_TYPE_CONTRACT_CALL_APPROVED),
		Chain:          parseAttributeValue(attribute["destination_chain"]).(string),
		//Todo: Fill ChainType by ChainID/ChainName
		ChainType:       "",
		Address:         address,
		ContractAddress: contractAddress,
		//Todo: Fill TransactionHash
		TransactionHash: "",
		ReturnValues: models.ReturnValuesDocument{
			SourceChain:     parseAttributeValue(attribute["chain"]).(string),
			SourceAddress:   parseAttributeValue(attribute["sender"]).(string),
			PayloadHash:     parseAttributeValue(attribute["payload_hash"]).(string),
			ContractAddress: parseAttributeValue(attribute["contract_address"]).(string),
			CommandID:       parseAttributeValue(attribute["command_id"]).(string),
		},
		Transaction: models.TransactionDocument{},
	}
	return approved
}

// Load data from event executed
func createExecuted(event *models.BlockEvent, attribute models.MapBlockEventAttributes) models.GMPStepDocument {
	eventId := parseAttributeValue(attribute["event_id"]).(string)
	executedChain := "" //Chain where message is executed
	executed := models.GMPStepDocument{
		ID:              eventId,
		Chain:           executedChain,
		SourceChain:     parseAttributeValue(attribute["source_chain"]).(string),
		ContractAddress: parseAttributeValue(attribute["contract_address"]).(string),
		Event:           normalizeEventType(event.BlockEventType.Type),
		TransactionHash: eventId,
		ReturnValues: models.ReturnValuesDocument{
			DestinationChain: parseAttributeValue(attribute["destination_chain"]).(string),
			PayloadHash:      parseAttributeValue(attribute["payload_hash"]).(string),
			Sender:           parseAttributeValue(attribute["sender"]).(string),
		},
	}
	return executed
}

func parseAttributeValue(value string) any {
	if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
		return value[1 : len(value)-1]
	}
	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		items := strings.Split(value[1:len(value)-1], ",")
		hexValue := "0x"
		for _, item := range items {
			v, e := strconv.Atoi(item)
			if e == nil {
				hexValue = fmt.Sprintf("%s%x", hexValue, v)
			}
		}
		return hexValue
	}
	return value
}
