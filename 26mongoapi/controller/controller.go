package controller

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://mustkeembaraskar:Pass2026@mongoapi.42w3gue.mongodb.net/?appName=mongoapi"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	// In v2, Connect instantiates and hooks up the client in a single operation
	client, err := mongo.Connect(clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo DB connection successful")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}
