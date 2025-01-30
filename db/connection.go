package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const (
	dbname = "go-todo"
)

func createConnection() (*mongo.Client, context.Context, context.CancelFunc) {

	credential := options.Credential{
		Username: "root",
		Password: "root",
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)

	client, err := mongo.Connect(clientOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, cancel
}

func closeConnection(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer func() {
		cancel()
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
