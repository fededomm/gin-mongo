package routes

import (
	"context"
	db "gin-mongo/database"
	"gin-mongo/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Routes struct {
	DB *mongo.Client
}

// GET ALL

func (r *Routes) GetOrdini(c *gin.Context) {

	var result []models.Ordini
	filter := bson.D{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	OrdiniCollection := db.GetCollection(r.DB, "Ordini")

	cursor, err := OrdiniCollection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messaggio": err.Error()})
		return
	}
	if err = cursor.All(ctx, &result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messaggio": err.Error()})
		return
	}
	if result == nil {
		c.JSON(http.StatusAccepted, gin.H{"messaggio": "nessun record"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "numero di record": len(result), "payload": result})
}

func (r *Routes) PostOrdini(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ordini := new(models.Ordini)
	defer cancel()
	if err := c.ShouldBindJSON(&ordini); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		log.Println(err.Error())
		return
	}

	postPayload := models.Ordini{
		NumeroOrdine: ordini.NumeroOrdine,
		Oggetto:      ordini.Oggetto,
		Data:         primitive.NewDateTimeFromTime(time.Now()),
		Destinatario: ordini.Destinatario,
		Mittente:     ordini.Mittente,
		Prezzo:       ordini.Prezzo,
	}

	ordiniCollection := db.GetCollection(r.DB, "Ordini")
	result, err := ordiniCollection.InsertOne(ctx, postPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	//metric.Counter.Add(ctx, 1)
	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "data": result, "payload": postPayload})
}
