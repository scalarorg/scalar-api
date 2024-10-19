package db

import (
	"context"

	"github.com/scalarorg/xchains-api/internal/db/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *Database) SaveDApp(ctx context.Context, dApp *model.DAppDocument) error {
	dApps := db.Client.Database(db.DbName).Collection(model.DAppCollection)
	// insert unique dApp
	_, err := dApps.InsertOne(ctx, dApp)
	return err
}

func (db *Database) GetDApp(ctx context.Context) ([]*model.DAppDocument, error) {
	dApps := db.Client.Database(db.DbName).Collection(model.DAppCollection)
	// get the cursor to iterator over the dApps
	cursor, err := dApps.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var dAppDocuments []*model.DAppDocument
	for cursor.Next(ctx) {
		var dApp model.DAppDocument
		err := cursor.Decode(&dApp)
		if err != nil {
			return nil, err
		}
		dAppDocuments = append(dAppDocuments, &dApp)
	}
	return dAppDocuments, nil
}

func (db *Database) UpdateDApp(ctx context.Context, dApp *model.DAppDocument) error {
	dApps := db.Client.Database(db.DbName).Collection(model.DAppCollection)
	// convert ID to objectID
	_id, err := primitive.ObjectIDFromHex(dApp.ID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{
		"chain_name":             dApp.ChainName,
		"btc_address_hex":        dApp.BTCAddressHex,
		"public_key_hex":         dApp.PublicKeyHex,
		"smart_contract_address": dApp.SmartContractAddress,
		"chain_id":               dApp.ChainID,
		"chain_endpoint":         dApp.ChainEndpoint,
		"rpc_url":                dApp.RPCUrl,
		"access_token":           dApp.AccessToken,
	}}

	updateResult := dApps.FindOneAndUpdate(ctx, filter, update)
	return updateResult.Err()
}

func (db *Database) ToggleDApp(ctx context.Context, ID string) error {
	dApps := db.Client.Database(db.DbName).Collection(model.DAppCollection)
	// convert ID to objectID
	_id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": _id}
	var result model.DAppDocument
	err = dApps.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return err
	}
	update := bson.M{"$set": bson.M{"state": !result.State}}
	updateResult := dApps.FindOneAndUpdate(ctx, filter, update)
	return updateResult.Err()
}

func (db *Database) DeleteDApp(ctx context.Context, ID string) error {
	dApps := db.Client.Database(db.DbName).Collection(model.DAppCollection)
	// convert ID to objectID
	_id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": _id}
	_, err = dApps.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
