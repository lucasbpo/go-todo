package db

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func Insert(collection string, data any) (bson.ObjectID, error) {
	client, ctx, cancel := createConnection()

	conn := client.Database(dbname).Collection(collection)

	res, err := conn.InsertOne(context.Background(), data)

	if err != nil {
		return bson.NilObjectID, err
	}

	defer closeConnection(client, ctx, cancel)
	return res.InsertedID.(bson.ObjectID), nil
}
