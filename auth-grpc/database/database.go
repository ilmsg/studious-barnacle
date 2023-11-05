package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase struct {
	// *mongo.Client
	*mongo.Database
}

func NewMongoDatabase(mongoURL, dbName string) *mongoDatabase {
	opts := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbName)
	return &mongoDatabase{db}
}

func (db *mongoDatabase) GetMongoCollection(collectionName string) *mongo.Collection {
	collection := db.Collection(collectionName)
	return collection
}
