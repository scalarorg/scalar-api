package services

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/postgres/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

func (s *Services) SearchBlocks(ctx context.Context, payload *types.SearchBlocksRequestPayload) ([]*types.SearchBlockResponsePayload, *types.Error) {
	blocks, err := s.DbAdapter.SearchBlocks(ctx, payload)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}
	if blocks == nil {
		blocks = []*models.Block{}
	}

	response := make([]*types.SearchBlockResponsePayload, len(blocks))
	for i, block := range blocks {
		response[i] = &types.SearchBlockResponsePayload{
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

func (s *Services) SearchBlockByHeight(ctx context.Context, height string) (*types.SearchBlockByHeightRequestPayload, *types.Error) {
	// Get the block from the database by height
	blocksFromDb, err := s.DbAdapter.GetBlocksByHeight(ctx, height)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}

	// Get the events from the database according to the block's ID
	eventsFromDb, err := s.DbAdapter.GetEventsByBlockID(ctx, blocksFromDb.ID)
	if err != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
	}

	// Get the events attributes from db according to the event ID
	eventIDs := []uint{}
	for _, event := range eventsFromDb {
		eventIDs = append(eventIDs, event.ID)
	}
	eventsAttributesFromDb, err := s.DbAdapter.GetEventsAttributesByEventIDs(ctx, eventIDs)
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
		BeginBlockEvents: beginBlockEvents,
		EndBlockEvents:   endBlockEvents,
	}
	return block, nil
}
