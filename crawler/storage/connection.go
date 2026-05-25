// Package storage
package storage

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConnection struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoConnection(dbName string) *MongoConnection {
	uri := os.Getenv("MONGO_DB_URL")
	if uri == "" {
		log.Fatalln("MONGO DB URI not found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln("ERROR Failed to connect to Mongo DB: ", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalln("ERROR Mongo DB Ping Failed: ", err)
	}

	log.Println("INFO Connected to MongoDB")

	return &MongoConnection{
		Client: client,
		DB:     client.Database(dbName),
	}
}

func (m *MongoConnection) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := m.Client.Disconnect(ctx); err != nil {
		log.Printf("Error disconnecting MongoDB: %v", err)
	}
}
