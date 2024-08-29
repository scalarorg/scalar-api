package gmp

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/scalarorg/xchains-api/internal/db/postgres"
	"github.com/scalarorg/xchains-api/internal/db/postgres/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

const (
	EVENT_TYPE_MESSAGE                  string = "message"
	EVENT_TYPE_CONTRACT_CALL            string = "ContractCall"
	EVENT_TYPE_CONTRACT_CALL_WITH_TOKEN string = "ContractCallWithToken"
	EVENT_TYPE_CONTRACT_CALL_APPROVED   string = "axelar.evm.v1beta1.ContractCallApproved"
	EVENT_TYPE_MESSAGE_PROCESSING       string = "axelar.nexus.v1beta1.MessageProcessing"
	EVENT_TYPE_MESSAGE_EXECUTED         string = "axelar.nexus.v1beta1.MessageExecuted"
	EVENT_TYPE_XCHAINS_CONFIRM          string = "xchains.confirm"
)

type GmpClient struct {
	indexer *postgres.IndexerClient
	relayer *postgres.RelayerClient
}

func New(indexer *postgres.PostgresClient, relayer *postgres.PostgresClient) *GmpClient {
	return &GmpClient{
		indexer: &postgres.IndexerClient{
			PgClient: indexer,
		},
		relayer: &postgres.RelayerClient{
			PgClient: relayer,
		},
	}
}
func normalizeEventType(eventType string) string {
	parts := strings.Split(eventType, ".")
	return parts[len(parts)-1]
}
func (c *GmpClient) GMPSearch(ctx context.Context, payload *types.GmpPayload) ([]*GMPDocument, *types.Error) {
	options := &postgres.Options{}
	if payload != nil {
		options.Size = payload.Size
		options.Offset = payload.From
		options.EventId = payload.MessageID
	}
	if payload != nil {
		if payload.MessageID != "" {
			return c.getGMPByMessageID(ctx, payload.MessageID, options)
		}
		if payload.TxHash != "" {
			return c.getGMPByTxHash(ctx, payload.TxHash, options)
		}
	}

	relayDatas, err := c.relayer.GetRelayerDatas(ctx, options)
	if err != nil {
		return nil, err
	}

	return c.getGMPByRelayDatas(ctx, relayDatas)
}
func (c *GmpClient) getGMPByMessageID(ctx context.Context, messageID string, options *postgres.Options) ([]*GMPDocument, *types.Error) {
	relayDatas, err := c.relayer.GetRelayerDatas(ctx, options)
	if err != nil {
		return nil, err
	}
	return c.getGMPByRelayDatas(ctx, relayDatas)
}
func (c *GmpClient) getGMPByTxHash(ctx context.Context, txHash string, options *postgres.Options) ([]*GMPDocument, *types.Error) {
	relayDatas, err := c.relayer.GetRelayerDatas(ctx, options)
	if err != nil {
		return nil, err
	}
	return c.getGMPByRelayDatas(ctx, relayDatas)
}

func (c *GmpClient) getGMPByRelayDatas(ctx context.Context, relayDatas []postgres.RelayData) ([]*GMPDocument, *types.Error) {
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
	mapGmps := make(map[string]*GMPDocument)
	mapRelayDatas := make(map[string]postgres.RelayData)
	gmps := make([]*GMPDocument, len(relayDatas))
	for index, relayData := range relayDatas {
		gmps[index] = &GMPDocument{
			ID: relayData.ID,
			// From:   relayData.From,
			// To:     relayData.To,
			Status: relayData.Status.String,
			// CreatedAt: relayData.CreatedAt,
			// UpdatedAt: relayData.UpdatedAt,
		}
		createContractCall(gmps[index], &relayData)
		createContractCallApproved(gmps[index], &relayData)
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
	// mapContractCalls, err := c.relayer.GetContractCallParams(ctx, messageIds)
	// if err == nil {
	// 	contractCall := mapContractCalls[debugMsg]
	// 	fmt.Printf("Contract call %v\n", contractCall)
	// 	// for _, gmp := range gmps {
	// 	// 	// gmp.Call.TransactionHash = contractCall.PayloadHash
	// 	// 	// if contractCall, ok := mapContractCalls[gmp.Call.ID]; ok {
	// 	// 	// 	fmt.Printf("Contract call %v", contractCall)
	// 	// 	// 	// gmp.Call.Params = contractCall.Params
	// 	// 	// }

	// 	// }
	// } else {
	// 	fmt.Printf("Get contract call with Error %v", err)
	// }
	return gmps, nil
}

func createContractCall(gmp *GMPDocument, relayData *postgres.RelayData) {
	parts := strings.Split(relayData.ID, "-")
	call := GMPStepDocument{
		ID:              relayData.ID,
		Chain:           relayData.From.String,
		Event:           normalizeEventType(EVENT_TYPE_CONTRACT_CALL),
		ContractAddress: relayData.ContractCall.ContractAddress.String,
		Transaction: TransactionDocument{
			Hash: parts[0],
			//Xchains Todo: This field fill to Sender on the UI
			//From: relayData.From.String,
		},
		ReturnValues: ReturnValuesDocument{
			DestinationChain: relayData.To.String,
		},
		BlockNumber:     uint64(relayData.ContractCall.BlockNumber.Int32),
		TransactionHash: parts[0],
		//XChains Todo: Created
		BlockTimestamp: relayData.CreatedAt.Time.Unix(),
	}
	if len(parts) > 1 {
		if index, err := strconv.Atoi(parts[1]); err == nil {
			call.LogIndex = uint(index)
		}
	}
	if relayData.ContractCall.ContractAddress.Valid {
		call.ContractAddress = relayData.ContractCall.ContractAddress.String
		call.ReturnValues.Sender = relayData.ContractCall.SourceAddress.String
		call.ReturnValues.PayloadHash = relayData.ContractCall.PayloadHash.String
		call.ReturnValues.Payload = relayData.ContractCall.Payload.String
		call.ReturnValues.SourceAddress = relayData.ContractCall.SourceAddress.String
		call.ReturnValues.DestinationContractAddress = relayData.ContractCall.ContractAddress.String
		call.ReturnValues.DestinationChain = relayData.To.String
		call.ReturnValues.ContractAddress = relayData.ContractCall.ContractAddress.String
	}
	if relayData.ContractCallWithToken.ContractAddress.Valid {
		call.ContractAddress = relayData.ContractCallWithToken.ContractAddress.String
		call.ReturnValues.Sender = relayData.ContractCallWithToken.SourceAddress.String
		call.ReturnValues.PayloadHash = relayData.ContractCallWithToken.PayloadHash.String
		call.ReturnValues.Payload = relayData.ContractCallWithToken.Payload.String
		call.ReturnValues.SourceAddress = relayData.ContractCallWithToken.SourceAddress.String
		call.ReturnValues.DestinationContractAddress = relayData.ContractCallWithToken.ContractAddress.String
		call.ReturnValues.DestinationChain = relayData.To.String
		call.ReturnValues.ContractAddress = relayData.ContractCallWithToken.ContractAddress.String
	}
	gmp.Call = call
	// fmt.Printf("ContractCall %v\n", relayData.ContractCall)
	// fmt.Printf("ContractCallWithToken %v\n", relayData.ContractCallWithToken)
	// fmt.Printf("Call chain %s\n", call.Chain)
}

func createContractCallApproved(gmp *GMPDocument, relayData *postgres.RelayData) {
	approved := GMPStepDocument{
		//Todo: Fetch blockhash/Number/Timestamp?
		BlockHash: "",
		// BlockTimestamp: 0,
		Event: normalizeEventType(EVENT_TYPE_CONTRACT_CALL_APPROVED),
		//Todo: Fill ChainType by ChainID/ChainName
		ChainType:       "",
		Address:         relayData.ContractCall.SourceAddress.String,
		ContractAddress: relayData.ContractCall.ContractAddress.String,
	}

	if relayData.ContractCall.ContractCallApproved.SourceTxHash.Valid {
		approved.ID = relayData.ContractCall.ContractCallApproved.ID
		approved.BlockNumber = uint64(relayData.ContractCall.ContractCallApproved.BlockNumber.Int32)
		// approved.BlockTimestamp = relayData.ContractCall.ContractCallApproved.CreatedAt.Time.Unix()
		//approved.BlockHash = relayData.ContractCall.ContractCallApproved.BlockHash.String
		approved.Chain = relayData.ContractCall.ContractCallApproved.SourceChain.String
		//approved.Address = relayData.ContractCall.ContractCallApproved.SourceAddress.String
		//approved.ContractAddress = relayData.ContractCall.ContractCallApproved.ContractAddress.String
		approved.TransactionHash = relayData.ContractCall.ContractCallApproved.TxHash.String
		approved.ReturnValues.SourceChain = relayData.ContractCall.ContractCallApproved.SourceChain.String
		approved.ReturnValues.SourceEventIndex = string(relayData.ContractCall.ContractCallApproved.SourceEventIndex.Int64)
		approved.ReturnValues.SourceTxHash = relayData.ContractCall.ContractCallApproved.SourceTxHash.String
		approved.ReturnValues.SourceAddress = relayData.ContractCall.ContractCallApproved.SourceAddress.String
		approved.ReturnValues.ContractAddress = relayData.ContractCall.ContractCallApproved.ContractAddress.String
		approved.ReturnValues.PayloadHash = relayData.ContractCall.ContractCallApproved.PayloadHash.String
		approved.ReturnValues.CommandID = relayData.ContractCall.ContractCallApproved.CommandId.String

	}
	gmp.Approved = approved
}
func createGMPDocument(gmp *GMPDocument, relayData *postgres.RelayData, event *models.BlockEvent, attribute postgres.MapBlockEventAttributes) {
	gmp.ID = parseAttributeValue(attribute["event_id"]).(string)
	gmp.CommandID = parseAttributeValue(attribute["command_id"]).(string)
	// fmt.Printf("Event id %s\n", attribute["event_id"])
	switch event.BlockEventType.Type {
	case EVENT_TYPE_CONTRACT_CALL_APPROVED:
		gmp.Approved = createApprovedEvent(relayData, event, attribute)
	}
	// gmp.Call = createContractCall(event, attribute)
	// gmp.Confirm = createConfirmEvent(event, attribute)
	// gmp.Approved = createApprovedEvent(event, attribute)
	// gmp.Executed = createExecuted(event, attribute)
}

func createConfirmEvent(event *models.BlockEvent, attribute postgres.MapBlockEventAttributes) ConfirmDocument {
	// eventId := parseAttributeValue(attribute["event_id"]).(string)
	confirm := ConfirmDocument{
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

// Load data from event ContractCallApproved
func createApprovedEvent(relayData *postgres.RelayData, event *models.BlockEvent, attribute postgres.MapBlockEventAttributes) GMPStepDocument {
	eventId := parseAttributeValue(attribute["event_id"]).(string)
	//Todo fill address and contract address
	address := ""
	contractAddress := "" // The same as address
	approved := GMPStepDocument{
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
		ReturnValues: ReturnValuesDocument{
			SourceChain:     parseAttributeValue(attribute["chain"]).(string),
			SourceAddress:   parseAttributeValue(attribute["sender"]).(string),
			PayloadHash:     parseAttributeValue(attribute["payload_hash"]).(string),
			ContractAddress: parseAttributeValue(attribute["contract_address"]).(string),
			CommandID:       parseAttributeValue(attribute["command_id"]).(string),
		},
		Transaction: TransactionDocument{},
	}
	return approved
}

// Load data from event executed
func createExecuted(event *models.BlockEvent, attribute postgres.MapBlockEventAttributes) GMPStepDocument {
	eventId := parseAttributeValue(attribute["event_id"]).(string)
	executedChain := "" //Chain where message is executed
	executed := GMPStepDocument{
		ID:              eventId,
		Chain:           executedChain,
		SourceChain:     parseAttributeValue(attribute["source_chain"]).(string),
		ContractAddress: parseAttributeValue(attribute["contract_address"]).(string),
		Event:           normalizeEventType(event.BlockEventType.Type),
		TransactionHash: eventId,
		ReturnValues: ReturnValuesDocument{
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
