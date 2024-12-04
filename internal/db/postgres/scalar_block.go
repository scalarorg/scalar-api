package postgres

import (
	"context"

	"github.com/scalarorg/xchains-api/internal/db/postgres/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

const (
	recentBlocksLimit = 100
)

func (db *DbAdapter) SearchBlocks(ctx context.Context, payload *types.SearchBlocksRequestPayload) ([]*models.Block, error) {
	var size int
	if payload != nil && payload.Size != 0 {
		size = payload.Size
	} else {
		size = recentBlocksLimit
	}

	var blocks []*models.Block
	db.XchainsIndexerClient.WithContext(ctx).
		Order("height DESC").
		Preload("ProposerConsAddress").
		Limit(size).
		Find(&blocks)
	return blocks, nil
}

func (db *DbAdapter) GetBlocksByHeight(ctx context.Context, height string) (*models.Block, error) {
	var block *models.Block
	db.XchainsIndexerClient.WithContext(ctx).
		Preload("ProposerConsAddress").
		Where("height = ?", height).
		First(&block)
	return block, nil
}

func (db *DbAdapter) GetEventsByBlockID(ctx context.Context, blockID uint) ([]*models.BlockEvent, error) {
	var events []*models.BlockEvent
	db.XchainsIndexerClient.WithContext(ctx).
		Preload("BlockEventType").
		Where("block_id = ?", blockID).
		Find(&events)
	return events, nil
}

func (db *DbAdapter) GetEventsAttributesByEventIDs(ctx context.Context, eventIDs []uint) ([]*models.BlockEventAttribute, error) {
	var eventsAttributes []*models.BlockEventAttribute
	db.XchainsIndexerClient.WithContext(ctx).
		Preload("BlockEventAttributeKey").
		Where("block_event_id IN (?)", eventIDs).
		Find(&eventsAttributes)
	return eventsAttributes, nil
}

func (db *DbAdapter) GetEventsAttributesByEventID(ctx context.Context, eventID uint) ([]*models.BlockEventAttribute, error) {
	var eventsAttributes []*models.BlockEventAttribute
	db.XchainsIndexerClient.WithContext(ctx).
		Preload("BlockEventAttributeKey").
		Where("block_event_id = ?", eventID).
		Find(&eventsAttributes)
	return eventsAttributes, nil
}
