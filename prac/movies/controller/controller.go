package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const mongo_url = "mongodb+srv://10ohekabanzai:f22pakfaamcA@cluster0.fmpgski.mongodb.net/"

const db = "movies"
const col = "watchlist"

var collection *mongo.Collection

func Init() {

	// Create a context with a timeout to use later for operations like ping.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options using the connection URI.
	clientOptions := options.Client().ApplyURI(mongo_url)

	// Connect to MongoDB (v2 driver: no context parameter here).
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Use the context when pinging to verify the connection.
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	// Assign the collection from the database.
	collection := client.Database(db).Collection(col)
	fmt.Println("Connected to MongoDB!")
	fmt.Println("Collection Name:", collection.Name())

}
