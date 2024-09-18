package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/xchains-api/internal/types"
)

type RelayerClient struct {
	PgClient *PostgresClient
}

const QUERY_RELAYDATA = `SELECT rd.id, rd.status, rd.from, rd.to, rd."packetSequence", rd."executeHash", rd."createdAt", rd."updatedAt",
				c."blockNumber" as c_blockNumber, c."contractAddress" as c_contractAddress, c.payload as c_payload, c."payloadHash" as c_payloadHash, c."sourceAddress" as c_sourceAddress, c."stakerPublicKey" as c_stakerPublicKey,
				ca."sourceChain" as ca_sourceChain, ca."destinationChain" as ca_destinationChain, ca."txHash" as ca_txHash, ca."blockNumber" as ca_blockNumber, ca."logIndex" as ca_logIndex, ca."sourceAddress" as ca_sourceAddress,
				ca."contractAddress" as ca_contractAddress, ca."sourceTxHash" as ca_sourceTxHash, ca."sourceEventIndex" as ca_sourceEventIndex, ca."payloadHash" as ca_payloadHash, ca."commandId" as ca_commandId,
				ct."contractAddress" as ct_contractAddress, ct.amount as ct_amount, ct.symbol as ct_symbol, ct.payload as ct_payload, ct."payloadHash" as ct_payloadHash, ct."sourceAddress" as ct_sourceAddress
				FROM "RelayData" as rd
					LEFT JOIN "CallContract" as c ON c.id = rd.id
					LEFT JOIN "CallContractApproved" as ca ON c."sourceAddress" = ca."sourceAddress" AND c."contractAddress" = ca."contractAddress" AND c."payloadHash" = ca."payloadHash"
					LEFT JOIN "CallContractWithToken" as ct ON ct.id = rd.id`

func (c *RelayerClient) GetRelayerDatas(ctx context.Context, options *Options) ([]RelayData, *types.Error) {
	var relayDatas []RelayData
	if options.Size <= 0 {
		options.Size = 10
	}
	if options.Offset < 0 {
		options.Offset = 0
	}
	query := QUERY_RELAYDATA
	log.Ctx(ctx).Debug().Msg(fmt.Sprintf("GetRelayerDatas with Event Id: %s", options.EventId))
	if options.EventId != "" {
		query = query + " WHERE rd.id = ?"
		// query = query + " WHERE LEFT(rd.id, 66) = ?" // because the id was formated with the tx hash + event index like: 0x123...-0
	}
	query = query + fmt.Sprintf(` ORDER by rd."createdAt" desc OFFSET %d LIMIT %d`, options.Offset, options.Size)
	var rows *sql.Rows
	var err error
	if options.EventId != "" {
		rows, err = c.PgClient.Db.Raw(query, options.EventId).Rows()
	} else {
		rows, err = c.PgClient.Db.Raw(query).Rows()
	}
	log.Ctx(ctx).Debug().Msg(fmt.Sprintf("Query: %+v", err))
	if err != nil {
		return relayDatas, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	defer rows.Close()
	for rows.Next() {
		relayData := RelayData{
			ContractCall: ContractCall{
				ContractCallApproved: ContractCallApproved{},
			},
			ContractCallWithToken: ContractCallWithToken{},
		}
		// ScanRows scan a row into user
		err := rows.Scan(&relayData.ID,
			&relayData.Status,
			&relayData.From,
			&relayData.To,
			&relayData.PacketSequence,
			&relayData.ExecuteHash,
			&relayData.CreatedAt,
			&relayData.UpdatedAt,
			&relayData.ContractCall.BlockNumber,
			&relayData.ContractCall.ContractAddress,
			&relayData.ContractCall.Payload,
			&relayData.ContractCall.PayloadHash,
			&relayData.ContractCall.SourceAddress,
			&relayData.ContractCall.StakerPublicKey,
			&relayData.ContractCall.ContractCallApproved.SourceChain,
			&relayData.ContractCall.ContractCallApproved.DestinationChain,
			&relayData.ContractCall.ContractCallApproved.TxHash,
			&relayData.ContractCall.ContractCallApproved.BlockNumber,
			&relayData.ContractCall.ContractCallApproved.LogIndex,
			&relayData.ContractCall.ContractCallApproved.SourceAddress,
			&relayData.ContractCall.ContractCallApproved.ContractAddress,
			&relayData.ContractCall.ContractCallApproved.SourceTxHash,
			&relayData.ContractCall.ContractCallApproved.SourceEventIndex,
			&relayData.ContractCall.ContractCallApproved.PayloadHash,
			&relayData.ContractCall.ContractCallApproved.CommandId,
			&relayData.ContractCallWithToken.ContractAddress,
			&relayData.ContractCallWithToken.Amount,
			&relayData.ContractCallWithToken.Symbol,
			&relayData.ContractCallWithToken.Payload,
			&relayData.ContractCallWithToken.PayloadHash,
			&relayData.ContractCallWithToken.SourceAddress,
		)
		if err != nil {
			fmt.Printf("Error while scanning rows %v", err)
		}
		relayData.ContractCall.ID = relayData.ID
		relayData.ContractCallWithToken.ID = relayData.ID
		relayDatas = append(relayDatas, relayData)
		// do something
	}

	// result := c.PgClient.Db.Limit(options.Size).Offset(options.Offset).Order("createdAt desc").Find(&relayDatas)
	// if result.Error != nil {
	// 	return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, result.Error)
	// }
	return relayDatas, nil
}

// func (c *RelayerClient) GetRelayerDatas0(ctx context.Context, options *Options) ([]RelayData, *types.Error) {
// 	var relayDatas []RelayData
// 	if options.Size <= 0 {
// 		options.Size = 10
// 	}
// 	if options.Offset < 0 {
// 		options.Offset = 0
// 	}
// 	result := c.PgClient.Db.Limit(options.Size).Offset(options.Offset).Order("createdAt desc").Find(&relayDatas)
// 	if result.Error != nil {
// 		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, result.Error)
// 	}

// 	return relayDatas, nil
// }

func (c *RelayerClient) GetContractCallParams(ctx context.Context, messageIds []string) (map[string]ContractCall, *types.Error) {
	mapContractCalls := make(map[string]ContractCall)
	var contractCalls []ContractCall
	result := c.PgClient.Db.Where("id IN ?", messageIds).Find(&contractCalls)
	if result.Error != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, result.Error)
	}
	for _, contractCall := range contractCalls {
		mapContractCalls[contractCall.ID] = contractCall
	}
	return mapContractCalls, nil
}
