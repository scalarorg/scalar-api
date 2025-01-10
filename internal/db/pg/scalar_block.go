package pg

import (
	"context"

	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

const (
	recentBlocksLimit = 100
)

func (c *PostgresClient) SearchBlocks(ctx context.Context, payload *types.SearchBlocksRequestPayload) ([]*models.Block, error) {
	var size int
	if payload != nil && payload.Size != 0 {
		size = payload.Size
	} else {
		size = recentBlocksLimit
	}

	var blocks []*models.Block
	c.DB.WithContext(ctx).
		Order("height DESC").
		Preload("ProposerConsAddress").
		Limit(size).
		Find(&blocks)
	return blocks, nil
}

func (c *PostgresClient) GetBlocksByHeight(ctx context.Context, height string) (*models.Block, error) {
	var block *models.Block
	c.DB.WithContext(ctx).
		Preload("ProposerConsAddress").
		Where("height = ?", height).
		First(&block)
	return block, nil
}

func (c *PostgresClient) GetEventsByBlockID(ctx context.Context, blockID uint) ([]*models.BlockEvent, error) {
	var events []*models.BlockEvent
	c.DB.WithContext(ctx).
		Preload("BlockEventType").
		Where("block_id = ?", blockID).
		Find(&events)
	return events, nil
}

func (c *PostgresClient) GetEventsAttributesByEventIDs(ctx context.Context, eventIDs []uint) ([]*models.BlockEventAttribute, error) {
	var eventsAttributes []*models.BlockEventAttribute
	c.DB.WithContext(ctx).
		Preload("BlockEventAttributeKey").
		Where("block_event_id IN (?)", eventIDs).
		Find(&eventsAttributes)
	return eventsAttributes, nil
}

func (c *PostgresClient) GetEventsAttributesByEventID(ctx context.Context, eventID uint) ([]*models.BlockEventAttribute, error) {
	var eventsAttributes []*models.BlockEventAttribute
	c.DB.WithContext(ctx).
		Preload("BlockEventAttributeKey").
		Where("block_event_id = ?", eventID).
		Find(&eventsAttributes)
	return eventsAttributes, nil
}
