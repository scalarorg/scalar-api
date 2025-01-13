package pg

import (
	"github.com/scalarorg/data-models/scalarnet"
	"github.com/scalarorg/xchains-api/internal/db/pg/models"
)

// func (c *PostgresClient) Search(ctx context.Context, payload *types.VaultPayload) ([]*models.VaultDocument, error) {
// 	options := &models.Options{}
// 	if payload != nil {
// 		// if payload.StakerPubkey != "" {
// 		// 	options.StakerPubkey = payload.StakerPubkey
// 		// }
// 	}

// 	relayerData, err := c.GetExecutedVaultBonding(ctx, options)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return c.getVaultsByRelayData(relayerData)
// }

func (c *PostgresClient) getVaultsByRelayData(callContracts []scalarnet.CallContract) ([]*models.VaultDocument, error) {
	vaults := make([]*models.VaultDocument, 0, len(callContracts))

	// TODO: why loop here? query batching?
	for _, callContract := range callContracts {
		// vault, err := c.getVaultByRelayData(&relayData)
		// if err != nil {
		// 	return nil, err
		// }

		_ = callContract
		vaults = append(vaults, &models.VaultDocument{})
	}
	return vaults, nil

}

// func (c *PostgresClient) getVaultByRelayData(vault *chains.Vault) (*models.VaultDocument, error) {
// 	// txHex := hex.EncodeToString(relayData.CallContract.TxHex)
// 	return &models.VaultDocument{
// 		// ID:                              relayData.ID,
// 		// Status:                          strconv.Itoa(int(relayData.Status)),
// 		// SimplifiedStatus:                string(models.ToReadableStatus(int(relayData.Status))),
// 		// SourceChain:                     relayData.From,
// 		// DestinationChain:                relayData.To,
// 		// DestinationSmartContractAddress: relayData.CallContract.DestContractAddress,
// 		// SourceTxHash:                    relayData.CallContract.TxHash,
// 		// SourceTxHex:                     txHex,
// 		// // Amount:                          relayData.CallContract.,
// 		// // StakerPubkey:   relayData.CallContract.StakerPublicKey,
// 		// CreatedAt: uint64(relayData.CreatedAt.Unix()),
// 		// UpdatedAt: uint64(relayData.UpdatedAt.Unix()),
// 		// ExecutedAmount: relayData.CallContract.CommandExecuted.Amount,
// 	}, nil
// }
