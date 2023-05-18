package database

import (
	"context"
	"gin-mongo/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitCounterCollection(db *mongo.Client) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	counterCollection := GetCollection(db, "counter")
	countDoc, err := counterCollection.CountDocuments(ctx, bson.M{})

	if countDoc == 0 {
		counterCollection.InsertOne(ctx, models.Counter{
			Id:  "ordineId",
			Seq: 1,
		})
	}

	if err != nil {
		log.Print("impossibile inizializzare la collection counter")
		log.Println("Nessuna connessione al database")
		return err
	}
	return nil 
}
