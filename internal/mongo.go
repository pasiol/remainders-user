package internal

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectOrDie(uri string, dbName string) (a *mongo.Database, b *mongo.Client) {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("fatal: %s", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("fatal: %s", err)
	}
	var DB = client.Database(dbName)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("fatal: %s", err)
	}

	return DB, client
}
