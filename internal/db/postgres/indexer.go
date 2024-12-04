package postgres

import (
	"context"
	"net/http"

	"github.com/scalarorg/xchains-api/internal/db/postgres/models"
	"github.com/scalarorg/xchains-api/internal/types"
)

type IndexerClient struct {
	PgClient *PostgresClient
}

type Options struct {
	Size         int
	Offset       int
	EventId      string
	EventType    string
	EventTypes   []string
	StakerPubkey string
}
type MapBlockEventAttributes map[string]string

func (c *IndexerClient) FindEventsByType(ctx context.Context, options *Options) ([]models.BlockEvent, *types.Error) {
	var events []models.BlockEvent
	// this can potentially be optimized by getting max first and selecting it (this gets translated into a select * limit 1)
	// c.Client.Joins("inner join txes on tx_messages.tx_id = txes.id").Find(&txMsgs)
	if options.Size <= 0 {
		options.Size = 10
	}
	if options.Offset < 0 {
		options.Offset = 0
	}
	result := c.PgClient.Db.Limit(options.Size).Offset(options.Offset).InnerJoins("BlockEventType", c.PgClient.Db.Where(&models.BlockEventType{Type: options.EventType})).Find(&events)
	if result.Error != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, result.Error)
	}

	return events, nil
}

// func (c *IndexerClient) FindEventsByTypes(ctx context.Context, options *Options) ([]models.BlockEvent, *types.Error) {
// 	var events []models.BlockEvent
// 	if options.Size <= 0 {
// 		options.Size = 10
// 	}
// 	if options.Offset < 0 {
// 		options.Offset = 0
// 	}
// 	if len(options.EventTypes) > 0 {
// 		query := `SELECT BlockEvent.id, block_id, block_event_type_id, lifecycle_position,
// 				BlockEventType.type as BlockEventType__type,
// 				Block.id as Block__id, Block.time_stamp as Block__time_stamp, Block.height as Block__height, Block.chain_id, Block.proposer_cons_address_id as Block__proposer_cons_address_id,
// 				Address.address as Block__proposer_address,
// 				Chain.chain_id as chain_id, Chain.name as Chain__name
// 			FROM block_events as BlockEvent
// 				JOIN block_event_types as BlockEventType ON BlockEventType.id = BlockEvent.block_event_type_id
// 				JOIN blocks as Block ON Block.id = BlockEvent.block_id
// 				JOIN chains as Chain ON Chain.id = Block.chain_id
// 				JOIN addresses as Address ON Address.id = Block.proposer_cons_address_id
// 			WHERE BlockEventType.Type IN ?`
// 		query = query + fmt.Sprintf(" OFFSET %d LIMIT %d", options.Offset, options.Size)
// 		rows, err := c.PgClient.Db.Raw(query, options.EventTypes).Rows()
// 		if err != nil {
// 			return events, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
// 		}
// 		defer rows.Close()

// 		for rows.Next() {
// 			blockEvent := models.BlockEvent{
// 				Block: models.Block{
// 					Chain:               models.Chain{},
// 					ProposerConsAddress: models.Address{},
// 				},
// 				BlockEventType: models.BlockEventType{},
// 			}
// 			// ScanRows scan a row into user
// 			rows.Scan(&blockEvent.ID,
// 				&blockEvent.BlockID,
// 				&blockEvent.BlockEventTypeID,
// 				&blockEvent.LifecyclePosition,
// 				&blockEvent.BlockEventType.Type,
// 				&blockEvent.Block.ID,
// 				&blockEvent.Block.TimeStamp,
// 				&blockEvent.Block.Height,
// 				&blockEvent.Block.ChainID,
// 				&blockEvent.Block.ProposerConsAddressID,
// 				&blockEvent.Block.ProposerConsAddress.Address,
// 				&blockEvent.Block.Chain.ChainID,
// 				&blockEvent.Block.Chain.Name,
// 			)
// 			blockEvent.BlockEventType.ID = blockEvent.BlockEventTypeID
// 			blockEvent.Block.ProposerConsAddress.ID = blockEvent.Block.ProposerConsAddressID
// 			blockEvent.Block.Chain.ID = blockEvent.Block.ChainID
// 			// fmt.Printf("BlockEvent %v", blockEvent)
// 			events = append(events, blockEvent)
// 			// do something
// 		}
// 	}
// 	return events, nil
// }

// func (c *IndexerClient) FindBlockEventsByAttribute(ctx context.Context, attribute string, values []string, options *Options) ([]models.BlockEvent, *types.Error) {
// 	var blockEvents []models.BlockEvent
// 	if options.Size <= 0 {
// 		options.Size = 32767
// 	}
// 	if options.Offset < 0 {
// 		options.Offset = 0
// 	}
// 	stringValues := make([]string, len(values))
// 	for index, value := range values {
// 		stringValues[index] = "\"" + value + "\""
// 	}
// 	// First get block event attributes
// 	query := `SELECT Attribute.block_event_id
// 			FROM block_event_attributes as Attribute JOIN block_event_attribute_keys as AttributeKey ON Attribute.block_event_attribute_key_id = AttributeKey.id
// 			WHERE AttributeKey.key = ? AND Attribute.value in ?`
// 	query = query + fmt.Sprintf(" OFFSET %d LIMIT %d", options.Offset, options.Size)
// 	rows, err := c.PgClient.Db.Raw(query, attribute, stringValues).Rows()
// 	if err != nil {
// 		return blockEvents, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
// 	}
// 	defer rows.Close()
// 	var eventIds []uint
// 	for rows.Next() {
// 		var eventId uint
// 		rows.Scan(&eventId)
// 		eventIds = append(eventIds, eventId)
// 	}
// 	if len(eventIds) > 0 {
// 		query := `SELECT BlockEvent.id, block_id, block_event_type_id, lifecycle_position,
// 				BlockEventType.type as BlockEventType__type,
// 				Block.id as Block__id, Block.time_stamp as Block__time_stamp, Block.height as Block__height, Block.chain_id, Block.proposer_cons_address_id as Block__proposer_cons_address_id,
// 				Address.address as Block__proposer_address,
// 				Chain.chain_id as chain_id, Chain.name as Chain__name
// 			FROM block_events as BlockEvent
// 				JOIN block_event_types as BlockEventType ON BlockEventType.id = BlockEvent.block_event_type_id
// 				JOIN blocks as Block ON Block.id = BlockEvent.block_id
// 				JOIN chains as Chain ON Chain.id = Block.chain_id
// 				JOIN addresses as Address ON Address.id = Block.proposer_cons_address_id
// 			WHERE BlockEvent.id IN ?`
// 		query = query + fmt.Sprintf(" OFFSET %d LIMIT %d", options.Offset, options.Size)
// 		rows, err := c.PgClient.Db.Raw(query, eventIds).Rows()
// 		if err != nil {
// 			return blockEvents, types.NewError(http.StatusInternalServerError, types.InternalServiceError, err)
// 		}
// 		defer rows.Close()

// 		for rows.Next() {
// 			blockEvent := models.BlockEvent{
// 				Block: models.Block{
// 					Chain:               models.Chain{},
// 					ProposerConsAddress: models.Address{},
// 				},
// 				BlockEventType: models.BlockEventType{},
// 			}
// 			// ScanRows scan a row into user
// 			rows.Scan(&blockEvent.ID,
// 				&blockEvent.BlockID,
// 				&blockEvent.BlockEventTypeID,
// 				&blockEvent.LifecyclePosition,
// 				&blockEvent.BlockEventType.Type,
// 				&blockEvent.Block.ID,
// 				&blockEvent.Block.TimeStamp,
// 				&blockEvent.Block.Height,
// 				&blockEvent.Block.ChainID,
// 				&blockEvent.Block.ProposerConsAddressID,
// 				&blockEvent.Block.ProposerConsAddress.Address,
// 				&blockEvent.Block.Chain.ChainID,
// 				&blockEvent.Block.Chain.Name,
// 			)
// 			blockEvent.BlockEventType.ID = blockEvent.BlockEventTypeID
// 			blockEvent.Block.ProposerConsAddress.ID = blockEvent.Block.ProposerConsAddressID
// 			blockEvent.Block.Chain.ID = blockEvent.Block.ChainID
// 			blockEvents = append(blockEvents, blockEvent)
// 		}
// 	}
// 	return blockEvents, nil
// }

func (c *IndexerClient) FindEventAttributes(ctx context.Context, eventIds []uint) (map[uint]MapBlockEventAttributes, *types.Error) {
	mapAttributes := make(map[uint]MapBlockEventAttributes)
	var attributes []models.BlockEventAttribute
	result := c.PgClient.Db.Preload("BlockEventAttributeKey").Where("block_event_id IN ?", eventIds).Find(&attributes)
	if result.Error != nil {
		return nil, types.NewError(http.StatusInternalServerError, types.InternalServiceError, result.Error)
	}
	for _, attribute := range attributes {
		if _, ok := mapAttributes[attribute.BlockEventID]; !ok {
			mapAttributes[attribute.BlockEventID] = make(MapBlockEventAttributes)
		}
		// fmt.Printf("Attribute id %d key %s => value %s\n", attribute.ID, attribute.BlockEventAttributeKey.Key, attribute.Value)
		mapAttributes[attribute.BlockEventID][attribute.BlockEventAttributeKey.Key] = attribute.Value
	}
	return mapAttributes, nil
}
