package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find the .env file")
	}

	MongoDB := os.Getenv("MONGODB_URL")

	if MongoDB == "" {
		log.Fatal("MONGODB_URL not set!")
	}
	fmt.Println("MongoDB URL", MongoDB)

	clientOptions := options.Client().ApplyURI(MongoDB) //mangodb client

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		return nil
	}

	return client

}

var Client *mongo.Client = DBInstance()

func OpenCollection(collectionName string, client *mongo.Client) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find the .env file")
	}
	databaseName := os.Getenv("DATABASE_NAME")
	fmt.Println("DATABASE_NAME:", databaseName)

	collection := client.Database("databaseName").Collection("collectionName")

	if collection == nil {
		return nil
	}
	return collection

}
