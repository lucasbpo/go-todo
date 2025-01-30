package db

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func Find(collection string, documents any) error {
	client, ctx, cancel := createConnection()

	coll := client.Database(dbname).Collection(collection)

	cursor, err := coll.Find(context.Background(), bson.D{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())
	defer closeConnection(client, ctx, cancel)
	return cursor.All(context.Background(), documents)
}

func FindByID(collection, id string, document any) error {
	client, ctx, cancel := createConnection()

	coll := client.Database(dbname).Collection(collection)

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	defer closeConnection(client, ctx, cancel)

	return coll.FindOne(context.Background(), filter).Decode(document)
}
