package postgres

import (
	"context"

	"github.com/scalarorg/xchains-api/internal/db/postgres/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

const (
	recentBlocksLimit = 100
)

func (db *DbAdapter) SearchBlocks(ctx context.Context, payload *types.SearchBlocksRequestPayload) ([]*models.Block, *types.Error) {
	var size int
	if payload != nil && payload.Size != 0 {
		size = payload.Size
	} else {
		size = recentBlocksLimit
	}

	var blocks []*models.Block
	db.XchainsIndexerClient.WithContext(ctx).
		Order("height DESC").
		Limit(size).
		Find(&blocks)
	return blocks, nil
}
