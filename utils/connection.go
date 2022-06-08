package utils

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE       = "url_shortner"
	CONNECTION_URL = "mongodb+srv://user:userpassword@mflix.uwgun.mongodb.net/?retryWrites=true&w=majority"
)

var DatabaseClient *mongo.Client

// ConnectDB will establish connection with mongodb
func ConnectDB() (client *mongo.Client) {
	// Set credentials
	// credentials := options.Credential{
	// 	Username: os.Getenv("MONGODB_USER"),
	// 	Password: os.Getenv("MONGODB_PASSWORD"),
	// }
	// Set client options
	clientOptions := options.Client().ApplyURI(CONNECTION_URL)

	// Connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	handlerError(err)
	log.Println("Connected to Mongodb")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) (*mongo.Collection, error) {
	db := client.Database(DATABASE)
	names, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	if !stringInSlice(collectionName, names) {
		return nil, errors.New("collection does not exist")
	}

	return db.Collection(collectionName), nil
}
