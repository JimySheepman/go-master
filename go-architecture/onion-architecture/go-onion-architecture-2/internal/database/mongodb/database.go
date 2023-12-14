// Package mongodb holds MongoDB database and repository methods.
package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const timeoutDefaultDuration = 5

// Database - MongoDB database struct
type Database struct {
	client   *mongo.Client
	database string
}

// New creates a new MongoDB database struct
func New(ctx context.Context, address string, database string) (*Database, error) {
	clientOptions := options.Client().ApplyURI(address)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDefaultDuration*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	db := Database{
		client:   client,
		database: database,
	}

	return &db, nil
}
