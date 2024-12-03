package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/postgres/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (s *Services) SearchBlocks(ctx context.Context, payload *types.SearchBlocksRequestPayload) ([]*types.SearchBlocksResponsePayload, *types.Error) {
	blocks, err := s.DbAdapter.SearchBlocks(ctx, payload)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	if blocks == nil {
		blocks = []*models.Block{}
	}

	response := make([]*types.SearchBlocksResponsePayload, len(blocks))
	for i, block := range blocks {
		response[i] = &types.SearchBlocksResponsePayload{
			ID:                    block.ID,
			Time:                  uint64(block.TimeStamp.Unix()),
			Height:                uint64(block.Height),
			ChainID:               block.ChainID,
			ProposerConsAddressID: block.ProposerConsAddressID,
			TxIndexed:             block.TxIndexed,
			BlockEventsIndexed:    block.BlockEventsIndexed,
		}
	}
	return response, nil
}
