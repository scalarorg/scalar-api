package db

import (
	"context"

	"github.com/babylonchain/staking-api-service/internal/db/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *Database) GetGMPs(ctx context.Context) ([]*model.GMPDocument, error) {
	gmps := db.Client.Database(db.DbName).Collection(model.GMPCollection)

	// get the cursor to iterator over the gmps
	cursor, err := gmps.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var gmpDocuments []*model.GMPDocument

	for cursor.Next(ctx) {
		var gmp model.GMPDocument
		err := cursor.Decode(&gmp)
		if err != nil {
			return nil, err
		}
		gmpDocuments = append(gmpDocuments, &gmp)
	}

	return gmpDocuments, nil
}
