package db

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func UpdateByID(collection, id string, data, result any) error {
	client, ctx, cancel := createConnection()

	coll := client.Database(dbname).Collection(collection)

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	opts := options.FindOneAndUpdate().SetUpsert(false)

	err = coll.FindOneAndUpdate(
		context.Background(), filter, bson.M{"$set": data}, opts).Err()
	if err != nil {
		return err
	}

	defer closeConnection(client, ctx, cancel)
	return coll.FindOne(context.Background(), filter).Decode(result)
}
