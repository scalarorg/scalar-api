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
    rd.id, rd.status, rd.from, rd.to, rd."packet_sequence", rd."execute_hash", rd."created_at", rd."updated_at",
    c.c_block_number, c.c_tx_hash,
    c.c_log_index, c.c_contract_address, c.c_payload, c.c_payload_hash, c.c_source_address, c.c_staker_public_key, c_sender_address, c.c_amount,
    ca.ca_source_chain, ca.ca_destination_chain, ca.ca_tx_hash, ca.ca_block_number, ca.ca_log_index, ca.ca_source_address,
    ca.ca_contract_address, ca.ca_source_tx_hash, ca.ca_source_event_index, ca.ca_payload_hash, ca.ca_command_id
FROM "relay_data" rd
LEFT JOIN (
    SELECT 
        c.id,
        c."block_number" as c_block_number, 
		c."tx_hash" as c_tx_hash,
		c."tx_hex" as c_tx_hex,
        c."log_index" as c_log_index, 
        c."contract_address" as c_contract_address, 
        c.payload as c_payload, 
        c."payload_hash" as c_payload_hash, 
        c."source_address" as c_source_address, 
        c."staker_public_key" as c_staker_public_key,
		c."sender_address" as c_sender_address,
		c."amount" as c_amount,
        ROW_NUMBER() OVER (PARTITION BY c.id ORDER BY c."block_number") as rn
    FROM "contract_calls" c
) c ON rd.id = c.id AND c.rn = 1
LEFT JOIN (
    SELECT 
        ca."source_address",
        ca."contract_address",
        ca."payload_hash",
        ca."source_chain" as ca_source_chain, 
        ca."destination_chain" as ca_destination_chain, 
        ca."tx_hash" as ca_tx_hash, 
        ca."block_number" as ca_block_number, 
        ca."log_index" as ca_log_index, 
        ca."source_address" as ca_source_address,
        ca."contract_address" as ca_contract_address, 
        ca."source_tx_hash" as ca_source_tx_hash, 
        ca."source_event_index" as ca_source_event_index, 
        ca."payload_hash" as ca_payload_hash, 
        ca."command_id" as ca_command_id,
        ROW_NUMBER() OVER (PARTITION BY ca."source_address", ca."contract_address", ca."payload_hash" ORDER BY ca."block_number") as rn
    FROM "contract_call_approveds" ca
) ca ON c.c_source_address = ca."source_address" AND c.c_contract_address = ca."contract_address" AND c.c_payload_hash = ca."payload_hash" AND ca.rn = 1`

// LEFT JOIN (
//     SELECT
//         ct.id,
//         ct."contract_address" as ct_contract_address,
//         ct.amount as ct_amount,
//         ct.symbol as ct_symbol,
//         ct.payload as ct_payload,
//         ct."payload_hash" as ct_payload_hash,
//         ct."source_address" as ct_source_address,
//         ROW_NUMBER() OVER (PARTITION BY ct.id ORDER BY ct."contract_address") as rn
//     FROM "call_contract_with_tokens" ct
// ) ct ON rd.id = ct.id AND ct.rn = 1`

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
	query = query + fmt.Sprintf(` ORDER by rd."created_at" desc OFFSET %d LIMIT %d`, options.Offset, options.Size)
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
			//ContractCallWithToken: ContractCallWithToken{},
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
			&relayData.ContractCall.SenderAddress,
			&relayData.ContractCall.Amount,
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
			// &relayData.ContractCallWithToken.ContractAddress,
			// &relayData.ContractCallWithToken.Amount,
			// &relayData.ContractCallWithToken.Symbol,
			// &relayData.ContractCallWithToken.Payload,
			// &relayData.ContractCallWithToken.PayloadHash,
			// &relayData.ContractCallWithToken.SourceAddress,
		)
		if err != nil {
			fmt.Printf("Error while scanning rows %v", err)
		}
		relayData.ContractCall.ID = relayData.ID
		//relayData.ContractCallWithToken.ID = relayData.ID
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
    rd.id, rd.status, rd.from, rd.to, rd."packet_sequence", rd."execute_hash", rd."created_at", rd."updated_at",
    c.c_block_number, c.c_tx_hash, c.c_tx_hex,
    c.c_log_index, c.c_contract_address, c.c_payload, c.c_payload_hash, c.c_source_address, c.c_staker_public_key, c.c_amount,
    ce.ce_amount
FROM "relay_data" rd
JOIN (
    SELECT 
        c.id,
        c."block_number" as c_block_number, 
        c."tx_hash" as c_tx_hash,
        c."tx_hex" as c_tx_hex,
        c."log_index" as c_log_index, 
        c."contract_address" as c_contract_address, 
        c.payload as c_payload, 
        c."payload_hash" as c_payload_hash, 
        c."source_address" as c_source_address, 
        c."staker_public_key" as c_staker_public_key,
        c."amount" as c_amount
    FROM "contract_calls" c
) c ON rd.id = c.id 
LEFT JOIN (
     SELECT ce."amount" as ce_amount,
            ce."reference_tx_hash" as ce_ref_tx_hash
     FROM "command_executeds" ce
) ce ON ce_ref_tx_hash = c_tx_hash WHERE rd."status" = '2'`

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
		query = query + " AND c.c_staker_public_key = ?"
	}

	query = query + fmt.Sprintf(` ORDER by rd."created_at" desc OFFSET %d LIMIT %d`, options.Offset, options.Size)

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
			//ContractCallWithToken: ContractCallWithToken{},
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
			&relayData.ContractCall.Amount,
			&relayData.ContractCall.CommandExecuted.Amount,
		)
		if err != nil {
			fmt.Printf("Error while scanning rows %v", err)
		}
		relayData.ContractCall.ID = relayData.ID
		//relayData.ContractCallWithToken.ID = relayData.ID
		relayDatas = append(relayDatas, relayData)
		// do something
	}

	// result := c.PgClient.Db.Limit(options.Size).Offset(options.Offset).Order("createdAt desc").Find(&relayDatas)
	// if result.Error != nil {
	// 	return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, result.Error)
	// }
	return relayDatas, nil
}