package pg

import (
	"context"
	"encoding/hex"
	"strconv"

	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (c *PostgresClient) Search(ctx context.Context, payload *types.VaultPayload) ([]*models.VaultDocument, error) {
	options := &models.Options{}
	if payload != nil {
		if payload.StakerPubkey != "" {
			options.StakerPubkey = payload.StakerPubkey
		}
	}

	relayerData, err := c.GetExecutedVaultBonding(ctx, options)
	if err != nil {
		return nil, err
	}

	return c.getVaultsByRelayData(relayerData)
}

func (c *PostgresClient) getVaultsByRelayData(relayDatas []models.RelayData) ([]*models.VaultDocument, error) {
	vaults := make([]*models.VaultDocument, 0, len(relayDatas))

	// TODO: why loop here? query batching?
	for _, relayData := range relayDatas {
		vault, err := c.getVaultByRelayData(&relayData)
		if err != nil {
			return nil, err
		}
		vaults = append(vaults, vault)
	}
	return vaults, nil

}

func (c *PostgresClient) getVaultByRelayData(relayData *models.RelayData) (*models.VaultDocument, error) {
	txHex := hex.EncodeToString(relayData.ContractCall.TxHex)
	return &models.VaultDocument{
		ID:                              relayData.ID,
		Status:                          strconv.Itoa(int(relayData.Status.Int32)),
		SimplifiedStatus:                string(models.ToReadableStatus(int(relayData.Status.Int32))),
		SourceChain:                     relayData.From.String,
		DestinationChain:                relayData.To.String,
		DestinationSmartContractAddress: relayData.ContractCall.ContractAddress.String,
		SourceTxHash:                    relayData.ContractCall.TxHash.String,
		SourceTxHex:                     txHex,
		Amount:                          relayData.ContractCall.Amount.String,
		StakerPubkey:                    relayData.ContractCall.StakerPublicKey.String,
		CreatedAt:                       uint64(relayData.CreatedAt.Time.Unix()),
		UpdatedAt:                       uint64(relayData.UpdatedAt.Time.Unix()),
		ExecutedAmount:                  relayData.ContractCall.CommandExecuted.Amount.String,
	}, nil
}
