package postgres

import (
	"context"

	"github.com/scalarorg/xchains-api/internal/db/postgres/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

const (
	searchTransactionsLimit = 100
)

func (db *DbAdapter) GetTransactionsByBlockHeight(ctx context.Context, height int64) ([]*models.Tx, error) {
	var blockId int64
	err := db.XchainsIndexerClient.WithContext(ctx).
		Model(&models.Block{}).
		Where("height = ?", height).
		Select("id").
		Scan(&blockId).Error
	if err != nil {
		return nil, err
	}

	var transactions []*models.Tx
	err = db.XchainsIndexerClient.WithContext(ctx).
		Model(&models.Tx{}).
		Where("block_id = ?", blockId).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (db *DbAdapter) GetTransactionByHash(ctx context.Context, hash string) (*models.Tx, error) {
	var transaction models.Tx
	err := db.XchainsIndexerClient.WithContext(ctx).
		Where("hash = ?", hash).
		Preload("Block").
		First(&transaction).Error
	return &transaction, err
}

func (db *DbAdapter) SearchTransactions(ctx context.Context, payload *types.SearchTransactionsRequestPayload) ([]*models.Tx, int, error) {
	var size int
	if payload != nil && payload.Size != 0 {
		size = payload.Size
	} else {
		size = searchTransactionsLimit
	}

	var transactions []*models.Tx
	var total int64

	// Count total rows
	if err := db.XchainsIndexerClient.WithContext(ctx).Model(&models.Tx{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Query transactions with ordering by block height
	err := db.XchainsIndexerClient.WithContext(ctx).
		Model(&models.Tx{}).
		Preload("Block").
		Joins("JOIN blocks ON txes.block_id = blocks.id").
		Order("blocks.height DESC").
		Order("txes.id DESC").
		Limit(size).
		Find(&transactions).Error
	if err != nil {
		return nil, 0, err
	}
	return transactions, int(total), nil
}

func (db *DbAdapter) GetNumTxsByBlockIDs(ctx context.Context, blockIDs []uint) (map[uint]int, error) {
	numTxs := make(map[uint]int)
	for _, blockID := range blockIDs {
		var count int64
		if err := db.XchainsIndexerClient.WithContext(ctx).Model(&models.Tx{}).Where("block_id = ?", blockID).Count(&count).Error; err != nil {
			return nil, err
		}
		numTxs[blockID] = int(count)
	}
	return numTxs, nil
}