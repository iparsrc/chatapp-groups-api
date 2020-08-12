package domain

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

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
	// check for indexes.
	CheckForIndexes()
}

func CheckForIndexes() {
	groupC := db.Collection("groups")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	cur, err := groupC.Indexes().List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	ctxx, cancelx := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelx()
	for cur.Next(ctxx) {
		var index bson.M
		if err := cur.Decode(&index); err != nil {
			log.Fatal(nil)
		}
		if index["name"] == "groupname" {
			return
		}
	}
	// At this point we are sure that groupname is not an index in the collection.
	CreateIndex("groupname", true)
}

func CreateIndex(name string, unique bool) {
	groupC := db.Collection("groups")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	ops := options.IndexOptions{
		Name:   &name,
		Unique: &unique,
	}
	m := mongo.IndexModel{
		Keys:    bson.D{{Key: name, Value: 1}},
		Options: &ops,
	}
	n, err := groupC.Indexes().CreateOne(ctx, m)
	if err != nil {
		log.Fatal(err)
	}
	if n != name {
		log.Fatal("index name is not groupname.")
	}
}
