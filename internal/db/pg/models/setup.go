package models

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	StatsLockCollection             = "stats_lock"
	OverallStatsCollection          = "overall_stats"
	FinalityProviderStatsCollection = "finality_providers_stats"
	StakerStatsCollection           = "staker_stats"
	DelegationCollection            = "delegations"
	TimeLockCollection              = "timelock_queue"
	UnbondingCollection             = "unbonding_queue"
	BtcInfoCollection               = "btc_info"
	UnprocessableMsgCollection      = "unprocessable_messages"
	DAppCollection                  = "dapps"
	GMPCollection                   = "gmps"
)

type index struct {
	Indexes map[string]int
	Unique  bool
}

var collections = map[string][]index{
	StatsLockCollection:             {{Indexes: map[string]int{}}},
	OverallStatsCollection:          {{Indexes: map[string]int{}}},
	FinalityProviderStatsCollection: {{Indexes: map[string]int{"active_tvl": -1}, Unique: false}},
	StakerStatsCollection:           {{Indexes: map[string]int{"active_tvl": -1}, Unique: false}},
	DelegationCollection: {
		{Indexes: map[string]int{"staker_pk_hex": 1, "staking_tx.start_height": -1}, Unique: false},
		{Indexes: map[string]int{"staker_btc_address.taproot_address": 1, "staking_tx.start_timestamp": -1}, Unique: false},
	},
	TimeLockCollection:         {{Indexes: map[string]int{"expire_height": 1}, Unique: false}},
	UnbondingCollection:        {{Indexes: map[string]int{"unbonding_tx_hash_hex": 1}, Unique: true}},
	UnprocessableMsgCollection: {{Indexes: map[string]int{}}},
	BtcInfoCollection:          {{Indexes: map[string]int{}}},
	DAppCollection:             {{Indexes: map[string]int{"chain_name": 1, "btc_address_hex": 1, "public_key_hex": 1, "state": 1}, Unique: true}},
	GMPCollection:              {{Indexes: map[string]int{}, Unique: true}},
}

// func Setup(ctx context.Context, cfg *config.Config) error {
// 	return nil
// }

func createCollection(ctx context.Context, database *mongo.Database, collectionName string) {
	// Check if the collection already exists.
	if _, err := database.Collection(collectionName).Indexes().CreateOne(ctx, mongo.IndexModel{}); err != nil {
		log.Debug().Msg(fmt.Sprintf("Collection maybe already exists: %s, skip the rest. info: %s", collectionName, err))
		return
	}

	// Create the collection.
	if err := database.CreateCollection(ctx, collectionName); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Failed to create collection: " + collectionName)
		return
	}

	log.Debug().Msg("Collection created successfully: " + collectionName)
}

func createIndex(ctx context.Context, database *mongo.Database, collectionName string, idx index) {
	if len(idx.Indexes) == 0 {
		return
	}

	indexKeys := bson.D{}
	for k, v := range idx.Indexes {
		indexKeys = append(indexKeys, bson.E{Key: k, Value: v})
	}

	index := mongo.IndexModel{
		Keys:    indexKeys,
		Options: options.Index().SetUnique(idx.Unique),
	}

	if _, err := database.Collection(collectionName).Indexes().CreateOne(ctx, index); err != nil {
		log.Debug().Msg(fmt.Sprintf("Failed to create index on collection '%s': %v", collectionName, err))
		return
	}

	log.Debug().Msg("Index created successfully on collection: " + collectionName)
}

// func ConnectToDBAndMigrate(dbConfig config.Database) (*gorm.DB, error) {
// 	database, err := db.PostgresDbConnect(dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.User, dbConfig.Password, strings.ToLower(dbConfig.LogLevel))
// 	if err != nil {
// 		config.Log.Fatal("Could not establish connection to the database", err)
// 	}

// 	sqldb, _ := database.DB()
// 	sqldb.SetMaxIdleConns(10)
// 	sqldb.SetMaxOpenConns(100)
// 	sqldb.SetConnMaxLifetime(time.Hour)

// 	err = db.MigrateModels(database)
// 	if err != nil {
// 		config.Log.Error("Error running DB migrations", err)
// 	}

// 	return database, err
// }
