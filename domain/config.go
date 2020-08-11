package domain

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func ConnectDB(uri string) {
	// Prepare to connect the database.
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err.Error())
	}
	// Connect to the database.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err.Error())
	}
	// db can be used in package level.
	db = client.Database("groups")
}
