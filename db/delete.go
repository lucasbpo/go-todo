package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func DeleteByID(collection, id string) error {
	client, ctx, cancel := createConnection()

	coll := client.Database(dbname).Collection(collection)

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	result, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return fmt.Errorf("%d documents deleted", result.DeletedCount)
	}

	defer closeConnection(client, ctx, cancel)
	return nil
}
