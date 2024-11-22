package vault

import (
	"context"
	"encoding/hex"
	"strconv"

	"github.com/scalarorg/xchains-api/internal/db/postgres"
	"github.com/scalarorg/xchains-api/internal/types"
)

type VaultClient struct {
	relayer *postgres.RelayerClient
}

func New(scalarPostgresClient *postgres.PostgresClient) *VaultClient {
	return &VaultClient{
		relayer: &postgres.RelayerClient{
			PgClient: scalarPostgresClient,
		},
	}
}

func (c *VaultClient) Search(ctx context.Context, payload *types.VaultPayload) ([]*VaultDocument, error) {
	options := &postgres.Options{}
	if payload != nil {
		if payload.StakerPubkey != "" {
			options.StakerPubkey = payload.StakerPubkey
		}
	}

	relayerData, err := c.relayer.GetExecutedVaultBonding(ctx, options)
	if err != nil {
		return nil, err
	}

	return c.getVaultsByRelayData(relayerData)
}

func (c *VaultClient) getVaultsByRelayData(relayDatas []postgres.RelayData) ([]*VaultDocument, error) {
	vaults := make([]*VaultDocument, 0, len(relayDatas))

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

func (c *VaultClient) getVaultByRelayData(relayData *postgres.RelayData) (*VaultDocument, error) {
	txHex := hex.EncodeToString(relayData.ContractCall.TxHex)
	return &VaultDocument{
		ID:                              relayData.ID,
		Status:                          strconv.Itoa(int(relayData.Status.Int32)),
		SimplifiedStatus:                string(postgres.ToReadableStatus(int(relayData.Status.Int32))),
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
