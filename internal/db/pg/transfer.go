package pg

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/scalarorg/data-models/relayer"
	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (c *PostgresClient) TokenSearchTransfers(ctx context.Context, options *models.Options) ([]*models.TransferDocument, int, *types.Error) {
	relayData, total, err := c.GetTokenSentRelayData(ctx, options)
	if err != nil {
		return nil, 0, err
	}

	result, err := c.getTransferByRelayData(ctx, relayData)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (c *PostgresClient) getTransferByRelayData(ctx context.Context, relayData []relayer.RelayData) ([]*models.TransferDocument, *types.Error) {
	messageIds := make([]string, len(relayData))
	for index, event := range relayData {
		messageIds[index] = event.ID
	}

	transfers := make([]*models.TransferDocument, len(relayData))
	for index, relayData := range relayData {
		createdAt := relayData.CreatedAt
		ms := createdAt.UnixNano() / int64(time.Millisecond)
		hour := createdAt.UnixNano() / int64(time.Hour)
		day := createdAt.UnixNano() / int64(time.Hour*24)
		week := createdAt.UnixNano() / int64(time.Hour*24*7)
		month := createdAt.UnixNano() / int64(time.Hour*24*30)
		quarter := createdAt.UnixNano() / int64(time.Hour*24*90)
		year := createdAt.UnixNano() / int64(time.Hour*24*365)
		transfers[index] = &models.TransferDocument{
			ID:               relayData.ID,
			Type:             "send_token", // TODO: change it
			Status:           strconv.Itoa(int(relayData.Status)),
			SimplifiedStatus: string(models.ToReadableStatus(int(relayData.Status))),
			TransferID:       uint(0), // TODO: change it
			Send: models.SendInfo{
				TxHash: relayData.TokenSent.TxHash,
				Height: relayData.TokenSent.BlockNumber,
				Status: "success",
				Type:   strings.Split(relayData.From, "|")[0],
				CreatedAt: models.TimeInfo{
					MS:      ms,
					Hour:    hour,
					Day:     day,
					Week:    week,
					Month:   month,
					Quarter: quarter,
					Year:    year,
				},
				SourceChain:              relayData.From,
				SenderAddress:            relayData.TokenSent.SourceAddress,
				RecipientAddress:         relayData.TokenSent.DestinationAddress,
				Denom:                    relayData.TokenSent.Symbol,
				Amount:                   float64(relayData.TokenSent.Amount),
				Value:                    float64(relayData.TokenSent.Amount),
				DestinationChain:         relayData.To,
				OriginalSourceChain:      relayData.From,
				Fee:                      5,
				FeeValue:                 float64(5.2),
				AmountReceived:           float64(relayData.TokenSent.Amount - 5),
				OriginalDestinationChain: relayData.To,
				InsufficientFee:          true,
			},
			Link: models.LinkInfo{
				ID:                       relayData.ID,
				Denom:                    relayData.TokenSent.Symbol,
				OriginalDestinationChain: relayData.To,
				Height:                   relayData.TokenSent.BlockNumber,
				TxHash:                   relayData.TokenSent.TxHash,
				CreatedAt: models.TimeInfo{
					MS:      createdAt.UnixNano() / int64(time.Millisecond),
					Hour:    hour,
					Day:     day,
					Week:    week,
					Month:   month,
					Quarter: quarter,
					Year:    year,
				},
				SourceChain:         relayData.From,
				SenderAddress:       relayData.TokenSent.SourceAddress,
				RecipientAddress:    relayData.TokenSent.DestinationAddress,
				DestinationChain:    relayData.To,
				OriginalSourceChain: relayData.From,
			},
			TimeSpent: models.TimeSpent{},
			// TODO: Query from appove to get vote, command, and confirm
			Command: models.CommandInfo{
				Chain:       relayData.From,
				CommandID:   "0000000000000000000000000000000000000000000000000000000000000000",
				LogIndex:    uint(relayData.TokenSent.LogIndex),
				BatchID:     "0000000000000000000000000000000000000000000000000000000000000000",
				BlockNumber: uint64(relayData.TokenSent.BlockNumber),
				CreatedAt: models.TimeInfo{
					MS:      createdAt.UnixNano() / int64(time.Millisecond),
					Hour:    hour,
					Day:     day,
					Week:    week,
					Month:   month,
					Quarter: quarter,
					Year:    year,
				},
				Executed:         true,
				BlockTimestamp:   int64(relayData.TokenSent.BlockNumber),
				TransactionIndex: uint(relayData.TokenSent.LogIndex),
				TransferID:       uint(0),
				TransactionHash:  relayData.TokenSent.TxHash,
			},
			Vote: models.VoteInfo{
				TransactionID: relayData.TokenSent.TxHash,
				PollID:        strconv.FormatUint(uint64(relayData.TokenSent.BlockNumber), 10),
				SourceChain:   relayData.From,
				CreatedAt: models.TimeInfo{
					MS:      createdAt.UnixNano() / int64(time.Millisecond),
					Hour:    hour,
					Day:     day,
					Week:    week,
					Month:   month,
					Quarter: quarter,
					Year:    year,
				},
				DestinationChain: relayData.To,
				Confirmation:     true,
				Type:             "Vote",
				Event:            "token_sent",
				TxHash:           relayData.TokenSent.TxHash,
				Height:           relayData.TokenSent.BlockNumber,
				Status:           "success",
				TransferID:       uint(0),
				Failed:           false,
				Success:          true,
			},
			Confirm: models.ConfirmInfo{
				Amount:         strconv.FormatUint(uint64(relayData.TokenSent.Amount), 10),
				SourceChain:    relayData.From,
				DepositAddress: relayData.TokenSent.DestinationAddress,
				CreatedAt: models.TimeInfo{
					MS:      createdAt.UnixNano() / int64(time.Millisecond),
					Hour:    hour,
					Day:     day,
					Week:    week,
					Month:   month,
					Quarter: quarter,
					Year:    year,
				},
				TxHash:           relayData.TokenSent.TxHash,
				Height:           relayData.TokenSent.BlockNumber,
				Status:           "success",
				TransferID:       uint(0),
				Denom:            relayData.TokenSent.Symbol,
				DestinationChain: relayData.To,
				Type:             "vote",
			},
		}
	}
	return transfers, nil
}
