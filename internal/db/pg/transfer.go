package pg

import (
	"context"
	"strconv"

	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (c *PostgresClient) TokenSearchTransfers(ctx context.Context, options *models.Options) ([]*models.TransferDocument, int, *types.Error) {
	tokenSents, total, err := c.ListTokenSents(ctx, options)
	if err != nil {
		return nil, 0, err
	}

	// TODO: join with token_approvals and commands, votes

	transfers, err := c.getTransferByRelayData(tokenSents)
	if err != nil {
		return nil, 0, err
	}

	return transfers, total, nil
}

func (c *PostgresClient) getTransferByRelayData(tokenSents []*chains.TokenSent) ([]*models.TransferDocument, *types.Error) {
	messageIds := make([]string, len(tokenSents))
	for index, sent := range tokenSents {
		messageIds[index] = sent.EventID
	}

	transfers := make([]*models.TransferDocument, len(tokenSents))
	for index, sent := range tokenSents {
		createdAt := sent.CreatedAt
		timeInfo := models.FormatTimeInfo(createdAt)

		var sourceChain string
		var sourceChainInfo chain.ChainInfo
		err := sourceChainInfo.FromString(sent.SourceChain)
		if err != nil {
			sourceChain = sent.SourceChain
		} else {
			sourceChain = chain.GetDisplayedName(sourceChainInfo)
		}

		var destinationChain string
		var destinationChainInfo chain.ChainInfo
		err = destinationChainInfo.FromString(sent.DestinationChain)
		if err != nil {
			destinationChain = sent.DestinationChain
		} else {
			destinationChain = chain.GetDisplayedName(destinationChainInfo)
		}

		transfers[index] = &models.TransferDocument{
			ID:               sent.EventID,
			Type:             models.TransferTypeSendToken, // TODO: change it
			Status:           string(sent.Status),
			SimplifiedStatus: string(sent.Status),
			TransferID:       uint(0), // TODO: change it
			Send: models.SendInfo{
				TxHash:                   sent.TxHash,
				Height:                   sent.BlockNumber,
				Status:                   string(sent.Status),
				Type:                     sourceChainInfo.ChainType.String(),
				CreatedAt:                timeInfo,
				SourceChain:              sourceChain,
				SenderAddress:            sent.SourceAddress,
				RecipientAddress:         sent.DestinationAddress,
				Denom:                    sent.Symbol,
				Amount:                   float64(sent.Amount),
				Value:                    float64(sent.Amount),
				DestinationChain:         destinationChain,
				OriginalSourceChain:      sourceChain,
				Fee:                      0,
				FeeValue:                 0,
				AmountReceived:           float64(sent.Amount),
				OriginalDestinationChain: destinationChain,
				InsufficientFee:          true,
			},
			Link: models.LinkInfo{
				ID:                       sent.EventID,
				Denom:                    sent.Symbol,
				OriginalDestinationChain: destinationChain,
				Height:                   sent.BlockNumber,
				TxHash:                   sent.TxHash,
				CreatedAt:                timeInfo,
				SourceChain:              sourceChain,
				SenderAddress:            sent.SourceAddress,
				RecipientAddress:         sent.DestinationAddress,
				DestinationChain:         destinationChain,
				OriginalSourceChain:      sourceChain,
			},
			TimeSpent: models.TimeSpent{},
			// TODO: Query from appove to get vote, command, and confirm
			Command: models.CommandInfo{
				Chain:            sourceChain,
				CommandID:        "",
				LogIndex:         uint(sent.LogIndex),
				BatchID:          "",
				BlockNumber:      uint64(sent.BlockNumber),
				CreatedAt:        timeInfo,
				Executed:         true,
				BlockTimestamp:   int64(sent.BlockNumber),
				TransactionIndex: uint(sent.LogIndex),
				TransferID:       uint(0),
				TransactionHash:  sent.TxHash,
			},
			Vote: models.VoteInfo{
				TransactionID:    sent.TxHash,
				PollID:           strconv.FormatUint(uint64(sent.BlockNumber), 10),
				SourceChain:      sourceChain,
				CreatedAt:        timeInfo,
				DestinationChain: destinationChain,
				Confirmation:     true,
				Type:             models.VoteTypeVote,
				Event:            models.VoteEventTokenSent,
				TxHash:           sent.TxHash,
				Height:           sent.BlockNumber,
				Status:           "success",
				TransferID:       uint(0),
				Success:          true,
			},
			Confirm: models.ConfirmInfo{
				Amount:           strconv.FormatUint(uint64(sent.Amount), 10),
				SourceChain:      sourceChain,
				DepositAddress:   sent.DestinationAddress,
				CreatedAt:        timeInfo,
				TxHash:           sent.TxHash,
				Height:           sent.BlockNumber,
				Status:           "success",
				TransferID:       uint(0),
				Denom:            sent.Symbol,
				DestinationChain: destinationChain,
				Type:             models.ConfirmTypeVote,
			},
		}
	}
	return transfers, nil
}
