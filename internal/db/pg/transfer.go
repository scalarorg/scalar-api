package pg

import (
	"context"
	"sort"
	"strconv"
	"time"

	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/data-models/chains"
	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (c *PostgresClient) TokenSearchTransfers(ctx context.Context, options *models.Options) ([]*models.TransferDocument, int, *types.Error) {
	tokenSents, n, err := c.ListTokenSents(ctx, options)
	if err != nil {
		return nil, 0, err
	}

	// We also query contract call with tokens when souce chain = evm and destination chain = bitcoin

	contractCallWithTokens, m, err := c.ListEvmToBTCTransfers(ctx, options)
	if err != nil {
		return nil, 0, err
	}

	total := n + m

	// TODO: join with token_approvals and commands, votes

	transfers, err := c.getTransferByRelayData(tokenSents, contractCallWithTokens)
	if err != nil {
		return nil, 0, err
	}

	return transfers, total, nil
}

func (c *PostgresClient) getTransferByRelayData(tokenSents []*chains.TokenSent, contractCallWithTokens []*chains.ContractCallWithToken) ([]*models.TransferDocument, *types.Error) {
	transfers := make([]*models.TransferDocument, 0, len(tokenSents)+len(contractCallWithTokens))

	for _, sent := range tokenSents {
		transfers = append(transfers, createTransferDocument(sent))
	}

	for _, sent := range contractCallWithTokens {
		transfers = append(transfers, createTransferDocument(sent))
	}

	// sort descending
	sort.Slice(transfers, func(i, j int) bool {
		return transfers[i].Send.CreatedAt.MS > transfers[j].Send.CreatedAt.MS
	})

	return transfers, nil
}

// TokenSender interface defines common methods for token sending entities
type TokenSender interface {
	GetEventID() string
	GetSourceChain() string
	GetDestinationChain() string
	GetTxHash() string
	GetBlockNumber() uint64
	GetSourceAddress() string
	GetDestinationAddress() string
	GetSymbol() string
	GetAmount() uint64
	GetStatus() string
	GetLogIndex() uint32
	GetCreatedAt() time.Time
	GetType() models.TransferType
}

func createTransferDocument(sent interface{}) *models.TransferDocument {
	var sender TokenSender

	switch v := sent.(type) {
	case *chains.TokenSent:
		sender = &tokenSentAdapter{v}
	case *chains.ContractCallWithToken:
		sender = &contractCallAdapter{v}
	default:
		return nil
	}

	createdAt := sender.GetCreatedAt()
	timeInfo := models.FormatTimeInfo(createdAt)

	var sourceChain string
	var sourceChainInfo chain.ChainInfo
	if err := sourceChainInfo.FromString(sender.GetSourceChain()); err != nil {
		sourceChain = sender.GetSourceChain()
	} else {
		sourceChain = chain.GetDisplayedName(sourceChainInfo)
	}

	var destinationChain string
	var destinationChainInfo chain.ChainInfo
	if err := destinationChainInfo.FromString(sender.GetDestinationChain()); err != nil {
		destinationChain = sender.GetDestinationChain()
	} else {
		destinationChain = chain.GetDisplayedName(destinationChainInfo)
	}

	return &models.TransferDocument{
		ID:               sender.GetEventID(),
		Type:             models.TransferType(sender.GetType()),
		Status:           sender.GetStatus(),
		SimplifiedStatus: sender.GetStatus(),
		TransferID:       uint(0),
		Send: models.SendInfo{
			TxHash:                   sender.GetTxHash(),
			SourceChain:              sourceChain,
			OriginalSourceChain:      sender.GetSourceChain(),
			DestinationChain:         destinationChain,
			OriginalDestinationChain: sender.GetDestinationChain(),
			Height:                   sender.GetBlockNumber(),
			Status:                   sender.GetStatus(),
			Type:                     sourceChainInfo.ChainType.String(),
			CreatedAt:                timeInfo,
			SenderAddress:            sender.GetSourceAddress(),
			RecipientAddress:         sender.GetDestinationAddress(),
			Denom:                    sender.GetSymbol(),
			Amount:                   float64(sender.GetAmount()),
			Value:                    float64(sender.GetAmount()),
			Fee:                      0,
			FeeValue:                 0,
			AmountReceived:           float64(sender.GetAmount()),
			InsufficientFee:          false,
		},
		Link: models.LinkInfo{
			ID:                       sender.GetEventID(),
			Denom:                    sender.GetSymbol(),
			SourceChain:              sender.GetSourceChain(),
			OriginalSourceChain:      sourceChain,
			DestinationChain:         sender.GetDestinationChain(),
			OriginalDestinationChain: destinationChain,
			Height:                   sender.GetBlockNumber(),
			TxHash:                   sender.GetTxHash(),
			CreatedAt:                timeInfo,
			SenderAddress:            sender.GetSourceAddress(),
			RecipientAddress:         sender.GetDestinationAddress(),
		},
		TimeSpent: models.TimeSpent{},
		// TODO: Query from appove to get vote, command, and confirm
		Command: models.CommandInfo{
			Chain:            sourceChain,
			CommandID:        "",
			LogIndex:         uint(sender.GetLogIndex()),
			BatchID:          "",
			BlockNumber:      uint64(sender.GetBlockNumber()),
			CreatedAt:        timeInfo,
			Executed:         true,
			BlockTimestamp:   int64(sender.GetBlockNumber()),
			TransactionIndex: uint(sender.GetLogIndex()),
			TransferID:       uint(0),
			TransactionHash:  sender.GetTxHash(),
		},
		Vote: models.VoteInfo{
			TransactionID:    sender.GetTxHash(),
			PollID:           strconv.FormatUint(uint64(sender.GetBlockNumber()), 10),
			SourceChain:      sourceChain,
			DestinationChain: destinationChain,
			CreatedAt:        timeInfo,
			Confirmation:     true,
			Type:             models.VoteTypeVote,
			Event:            models.VoteEventTokenSent,
			TxHash:           sender.GetTxHash(),
			Height:           sender.GetBlockNumber(),
			Status:           "success",
			TransferID:       uint(0),
			Success:          true,
		},
		Confirm: models.ConfirmInfo{
			Amount:           strconv.FormatUint(uint64(sender.GetAmount()), 10),
			SourceChain:      sourceChain,
			DestinationChain: destinationChain,
			DepositAddress:   sender.GetDestinationAddress(),
			CreatedAt:        timeInfo,
			TxHash:           sender.GetTxHash(),
			Height:           sender.GetBlockNumber(),
			Status:           "success",
			TransferID:       uint(0),
			Denom:            sender.GetSymbol(),
			Type:             models.ConfirmTypeVote,
		},
	}
}

// Adapter for TokenSent
type tokenSentAdapter struct {
	*chains.TokenSent
}

// Implement TokenSender interface for TokenSent
func (t *tokenSentAdapter) GetEventID() string            { return t.EventID }
func (t *tokenSentAdapter) GetSourceChain() string        { return t.SourceChain }
func (t *tokenSentAdapter) GetDestinationChain() string   { return t.DestinationChain }
func (t *tokenSentAdapter) GetTxHash() string             { return t.TxHash }
func (t *tokenSentAdapter) GetBlockNumber() uint64        { return t.BlockNumber }
func (t *tokenSentAdapter) GetSourceAddress() string      { return t.SourceAddress }
func (t *tokenSentAdapter) GetDestinationAddress() string { return t.DestinationAddress }
func (t *tokenSentAdapter) GetSymbol() string             { return t.Symbol }
func (t *tokenSentAdapter) GetAmount() uint64             { return t.Amount }
func (t *tokenSentAdapter) GetStatus() string             { return string(t.Status) }
func (t *tokenSentAdapter) GetLogIndex() uint32           { return uint32(t.LogIndex) }
func (t *tokenSentAdapter) GetCreatedAt() time.Time       { return t.CreatedAt }
func (t *tokenSentAdapter) GetType() models.TransferType  { return models.TransferTypeSendToken }

// Adapter for ContractCallWithToken
type contractCallAdapter struct {
	*chains.ContractCallWithToken
}

// Implement TokenSender interface for ContractCallWithToken
func (c *contractCallAdapter) GetEventID() string            { return c.EventID }
func (c *contractCallAdapter) GetSourceChain() string        { return c.SourceChain }
func (c *contractCallAdapter) GetDestinationChain() string   { return c.DestinationChain }
func (c *contractCallAdapter) GetTxHash() string             { return c.TxHash }
func (c *contractCallAdapter) GetBlockNumber() uint64        { return c.BlockNumber }
func (c *contractCallAdapter) GetSourceAddress() string      { return c.SourceAddress }
func (c *contractCallAdapter) GetDestinationAddress() string { return c.DestinationAddress }
func (c *contractCallAdapter) GetSymbol() string             { return c.Symbol }
func (c *contractCallAdapter) GetAmount() uint64             { return c.Amount }
func (c *contractCallAdapter) GetStatus() string             { return string(c.Status) }
func (c *contractCallAdapter) GetLogIndex() uint32           { return uint32(c.LogIndex) }
func (c *contractCallAdapter) GetCreatedAt() time.Time       { return c.CreatedAt }
func (c *contractCallAdapter) GetType() models.TransferType  { return models.TransferTypeSendToken }
