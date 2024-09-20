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

const QUERY_RELAYDATA = `
SELECT 
    rd.id, rd.status, rd.from, rd.to, rd."packetSequence", rd."executeHash", rd."createdAt", rd."updatedAt",
    c.c_blockNumber, c.c_txHash,
    c.c_logIndex, c.c_contractAddress, c.c_payload, c.c_payloadHash, c.c_sourceAddress, c.c_stakerPublicKey,
    ca.ca_sourceChain, ca.ca_destinationChain, ca.ca_txHash, ca.ca_blockNumber, ca.ca_logIndex, ca.ca_sourceAddress,
    ca.ca_contractAddress, ca.ca_sourceTxHash, ca.ca_sourceEventIndex, ca.ca_payloadHash, ca.ca_commandId,
    ct.ct_contractAddress, ct.ct_amount, ct.ct_symbol, ct.ct_payload, ct.ct_payloadHash, ct.ct_sourceAddress
FROM "RelayData" rd
LEFT JOIN (
    SELECT 
        c.id,
        c."blockNumber" as c_blockNumber, 
		c."txHash" as c_txHash,
		c."txHex" as c_txHex,
        c."logIndex" as c_logIndex, 
        c."contractAddress" as c_contractAddress, 
        c.payload as c_payload, 
        c."payloadHash" as c_payloadHash, 
        c."sourceAddress" as c_sourceAddress, 
        c."stakerPublicKey" as c_stakerPublicKey,
        ROW_NUMBER() OVER (PARTITION BY c.id ORDER BY c."blockNumber") as rn
    FROM "CallContract" c
) c ON rd.id = c.id AND c.rn = 1
LEFT JOIN (
    SELECT 
        ca."sourceAddress",
        ca."contractAddress",
        ca."payloadHash",
        ca."sourceChain" as ca_sourceChain, 
        ca."destinationChain" as ca_destinationChain, 
        ca."txHash" as ca_txHash, 
        ca."blockNumber" as ca_blockNumber, 
        ca."logIndex" as ca_logIndex, 
        ca."sourceAddress" as ca_sourceAddress,
        ca."contractAddress" as ca_contractAddress, 
        ca."sourceTxHash" as ca_sourceTxHash, 
        ca."sourceEventIndex" as ca_sourceEventIndex, 
        ca."payloadHash" as ca_payloadHash, 
        ca."commandId" as ca_commandId,
        ROW_NUMBER() OVER (PARTITION BY ca."sourceAddress", ca."contractAddress", ca."payloadHash" ORDER BY ca."blockNumber") as rn
    FROM "CallContractApproved" ca
) ca ON c.c_sourceAddress = ca."sourceAddress" AND c.c_contractAddress = ca."contractAddress" AND c.c_payloadHash = ca."payloadHash" AND ca.rn = 1
LEFT JOIN (
    SELECT 
        ct.id,
        ct."contractAddress" as ct_contractAddress, 
        ct.amount as ct_amount, 
        ct.symbol as ct_symbol, 
        ct.payload as ct_payload, 
        ct."payloadHash" as ct_payloadHash, 
        ct."sourceAddress" as ct_sourceAddress,
        ROW_NUMBER() OVER (PARTITION BY ct.id ORDER BY ct."contractAddress") as rn
    FROM "CallContractWithToken" ct
) ct ON rd.id = ct.id AND ct.rn = 1`

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
			&relayData.ContractCall.TxHash,
			&relayData.ContractCall.LogIndex,
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

// SELECT
// *
// FROM "RelayData" rd
// JOIN (
//     SELECT *
//     FROM "CallContract" c
// ) c ON rd.id = c.id

const QUERY_VAULT_RELAYDATA = `
SELECT 
    rd.id, rd.status, rd.from, rd.to, rd."packetSequence", rd."executeHash", rd."createdAt", rd."updatedAt",
    c.c_blockNumber, c.c_txHash, c.c_txHex,
    c.c_logIndex, c.c_contractAddress, c.c_payload, c.c_payloadHash, c.c_sourceAddress, c.c_stakerPublicKey
FROM "RelayData" rd
JOIN (
    SELECT 
        c.id,
        c."blockNumber" as c_blockNumber, 
		c."txHash" as c_txHash,
		c."txHex" as c_txHex,
        c."logIndex" as c_logIndex, 
        c."contractAddress" as c_contractAddress, 
        c.payload as c_payload, 
        c."payloadHash" as c_payloadHash, 
        c."sourceAddress" as c_sourceAddress, 
        c."stakerPublicKey" as c_stakerPublicKey
    FROM "CallContract" c
) c ON rd.id = c.id WHERE rd."status" = 2
`

func (c *RelayerClient) GetExecutedVaultBonding(ctx context.Context, options *Options) ([]RelayData, *types.Error) {
	var relayDatas []RelayData
	if options.Size <= 0 {
		options.Size = 10
	}
	if options.Offset < 0 {
		options.Offset = 0
	}
	query := QUERY_VAULT_RELAYDATA
	if options.StakerPubkey != "" {
		query = query + " AND c.c_stakerPublicKey = ?"
	}

	query = query + fmt.Sprintf(` ORDER by rd."createdAt" desc OFFSET %d LIMIT %d`, options.Offset, options.Size)

	var rows *sql.Rows
	var err error
	if options.StakerPubkey != "" {
		rows, err = c.PgClient.Db.Raw(query, options.StakerPubkey).Rows()
	} else {
		rows, err = c.PgClient.Db.Raw(query).Rows()
	}
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
			&relayData.ContractCall.TxHash,
			&relayData.ContractCall.TxHex,
			&relayData.ContractCall.LogIndex,
			&relayData.ContractCall.ContractAddress,
			&relayData.ContractCall.Payload,
			&relayData.ContractCall.PayloadHash,
			&relayData.ContractCall.SourceAddress,
			&relayData.ContractCall.StakerPublicKey,
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
