package database

import (
	"context"
	"fmt"
	"gin-mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitCounterCollection(db *mongo.Client) {

	ctx := context.Background()
	counterCollection := GetCollection(db, "counter")
	countDoc, err := counterCollection.CountDocuments(ctx, bson.M{})

	if countDoc == 0 {
		counterCollection.InsertOne(ctx, models.Counter{
			Id:  "ordineId",
			Seq: 1,
		})
	}

	if err != nil {
		fmt.Print("impossibile inizializzare la collection counter")
		return
	}
}
