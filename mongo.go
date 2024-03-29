package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	//"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func MongoInit() {
	mongoURI := "mongodb://127.0.0.1:27017" //os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	databaseName := "mongo-weather-db" //os.Getenv("MONGO_DB_NAME")
	if databaseName == "" {
		log.Fatal("MONGO_DB_NAME environment variable is not set")
	}

	collectionName := "mongoWeather" //os.Getenv("MONGO_COLLECTION_NAME")
	if collectionName == "" {
		log.Fatal("MONGO_COLLECTION_NAME environment variable is not set")
	}

	collection = client.Database(databaseName).Collection(collectionName)
}

func MongoInsert(input TransformedData) {
	result, err := collection.InsertOne(context.Background(), input)
	if err != nil {
		log.Fatal(err)
	}

	id := result.InsertedID
	fmt.Println(id)
}

func MongoGetByCity(input TransformedData) {
	var result bson.M
	// hard code for testing purposes:
	city := "Senoia"
	err := collection.FindOne(context.TODO(), bson.D{{Key: "City", Value: city}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", city)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

func MongoUpdate() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// limited use case for testing:
	result := collection.FindOneAndUpdate(ctx, bson.D{{Key: "currentTime", Value: "2024-01-25 10:13"}},
		bson.D{{Key: "City", Value: "UPDATED_CITY"}})
	if result.Err() != nil {
		log.Fatal(result.Err())
	}

	doc := bson.M{}
	decodeErr := result.Decode(&doc)

	fmt.Println(doc, decodeErr)
}

func MongoDeleteItem(input TransformedData) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// also limited for testing:
	filter := bson.M{"currentTime": "2024-01-25 10:13"}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
		return
	}

	if result.DeletedCount == 0 {
		fmt.Println("Nothing deleted...")
		return
	}

	fmt.Println()
}
