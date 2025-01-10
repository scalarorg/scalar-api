package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/pg/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (s *Services) SearchBlocks(ctx context.Context, payload *types.SearchBlocksRequestPayload) ([]*types.SearchBlockResponsePayload, *types.Error) {
	blocks, err := s.Pg.SearchBlocks(ctx, payload)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	if blocks == nil {
		blocks = []*models.Block{}
	}

	// Search number of txs in each block
	blockIDs := []uint{}
	for _, block := range blocks {
		blockIDs = append(blockIDs, block.ID)
	}
	numTxs, err := s.Pg.GetNumTxsByBlockIDs(ctx, blockIDs)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}

	response := make([]*types.SearchBlockResponsePayload, len(blocks))
	for i, block := range blocks {
		response[i] = &types.SearchBlockResponsePayload{
			ID:              block.ID,
			Hash:            block.Hash,
			Time:            uint64(block.TimeStamp.Unix()),
			Height:          uint64(block.Height),
			ProposerAddress: block.ProposerConsAddress.Address,
			NumTxs:          numTxs[block.ID],
		}
	}
	return response, nil
}

func (s *Services) SearchBlockByHeight(ctx context.Context, height string) (*types.SearchBlockByHeightRequestPayload, *types.Error) {
	// Get the block from the database by height
	blocksFromDb, err := s.Pg.GetBlocksByHeight(ctx, height)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}

	// Get the number of txs in the block
	numTxs, err := s.Pg.GetNumTxsByBlockIDs(ctx, []uint{blocksFromDb.ID})
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}

	// Get the events from the database according to the block's ID
	eventsFromDb, err := s.Pg.GetEventsByBlockID(ctx, blocksFromDb.ID)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}

	// Get the events attributes from db according to the event ID
	eventIDs := []uint{}
	for _, event := range eventsFromDb {
		eventIDs = append(eventIDs, event.ID)
	}
	eventsAttributesFromDb, err := s.Pg.GetEventsAttributesByEventIDs(ctx, eventIDs)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	attributesMap := make(map[uint][]*models.BlockEventAttribute)
	for _, attr := range eventsAttributesFromDb {
		attributesMap[attr.BlockEventID] = append(attributesMap[attr.BlockEventID], attr)
	}

	// Create return object
	beginBlockEvents := []types.BlockEvent{}
	endBlockEvents := []types.BlockEvent{}
	for _, event := range eventsFromDb {
		toAdd := types.BlockEvent{
			Type:       event.BlockEventType.Type,
			Attributes: []types.Attribute{},
		}
		for _, attr := range attributesMap[event.ID] {
			toAdd.Attributes = append(toAdd.Attributes, types.Attribute{
				Key:   attr.BlockEventAttributeKey.Key,
				Value: attr.Value,
			})
		}
		if event.LifecyclePosition == models.BeginBlockEvent {
			beginBlockEvents = append(beginBlockEvents, toAdd)
		} else {
			endBlockEvents = append(endBlockEvents, toAdd)
		}
	}
	block := &types.SearchBlockByHeightRequestPayload{
		Height:           height,
		Hash:             blocksFromDb.Hash,
		ProposerAddress:  blocksFromDb.ProposerConsAddress.Address,
		Time:             uint64(blocksFromDb.TimeStamp.Unix()),
		NumTxs:           numTxs[blocksFromDb.ID],
		BeginBlockEvents: beginBlockEvents,
		EndBlockEvents:   endBlockEvents,
	}
	return block, nil
}
