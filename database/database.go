package database

import (
	"context"
	"log"
	"job-portal-backend/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	log.Println("Starting MongoDB connection...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	Client = client
	log.Println("MongoDB connected")
}

// GetCollection fetches the collection from the database.
func GetCollection(collectionName string) *mongo.Collection {
	if Client == nil {
		log.Fatal("MongoDB client is not initialized. Please call Connect() first.")
	}
	return Client.Database("jobportal").Collection(collectionName) // Replace with your actual DB name
}
