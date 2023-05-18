package database

import (
	"context"
	"fmt"
	database "gin-mongo/configuration"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(cfg *database.DbConfig) *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.ConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongoDB")
	return client

}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Gestionale").Collection(collectionName)
	return collection
}
