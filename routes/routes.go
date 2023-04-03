package routes

import (
	"context"
	db "gin-mongo/database"
	"gin-mongo/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Routes struct {
	DB *mongo.Client
}

// GetAllOrdini
//
//	@Summary		List All Ordini
//	@Description	Ritorna tutti gli ordini contenuti nella Collection
//	@Tags			Ordini
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/gest [get]
func (r *Routes) GetOrdini(c *gin.Context) {

	var ordini []models.Ordini
	filter := bson.D{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ordiniCollection := db.GetCollection(r.DB, "Ordini")

	cursor, err := ordiniCollection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messaggio": err.Error()})
		return
	}
	if err = cursor.All(ctx, &ordini); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messaggio": err.Error()})
		return
	}
	if ordini == nil {
		c.JSON(http.StatusAccepted, gin.H{"messaggio": "nessun record"})
		return
	}

	log.Println(len(ordini))
	c.JSON(http.StatusOK, ordini)
}

// GET one Ordine
//
//	@Summary		GET one Ordine
//	@Description	GET un record nella Collection Ordini
//	@Tags			Ordini
//	@Param			numeroOrdine	path	string	true	"Numero Ordine dell'Ordine"
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/gest/{numeroOrdine} [get]
func (r *Routes) GetSingleOrdine(c *gin.Context) {
	var ordini models.Ordini
	numOrdine := c.Param("numeroOrdine")
	numOrdinetoInt, rerr := strconv.Atoi(numOrdine)
	if rerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "STRONZO"})
		return
	}
	filter := bson.M{"numeroOrdine": numOrdinetoInt}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ordiniCollection := db.GetCollection(r.DB, "Ordini")

	err := ordiniCollection.FindOne(ctx, filter).Decode(&ordini)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "trovato", "payload": ordini})
}

// CreateRecord
//
//	@Summary		Post one Ordine
//	@Description	Crea un record nella Collection Ordini
//	@Tags			Ordini
//	@Accept			json
//	@Produce		json
//	@Param			json	body	string	true	"Inserisci un Ordine"
//	@Success		200
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/gest [post]
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
	//numeroOrdineCollection := db.GetCollection(r.DB, "NumeroOrdine")

	result, err := ordiniCollection.InsertOne(ctx, postPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "data": result, "payload": postPayload})
}

// Update Ordini
//
//	@Summary		Update one Ordine
//	@Description	Esegui l'update di un Ordine
//	@Tags			Ordini
//	@Param			numeroOrdine	path	string	true	"Numero Ordine dell'Ordine"
//	@Param			json	body	string	true	"Modifica un Ordine"
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/gest/{numeroOrdine} [put]
func (r *Routes) UpdateOrdine(c *gin.Context) {

	var ordini models.Ordini
	numOrdine := c.Param("numeroOrdine")
	numOrdinetoInt, rerr := strconv.Atoi(numOrdine)
	if rerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "STRONZO"})
		return
	}
	filter := bson.M{"numeroOrdine": numOrdinetoInt}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := c.ShouldBindJSON(&ordini)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messaggio": err.Error()})
		log.Println(err.Error())
		return
	}

	ordiniUpdate := bson.M{
		//"numeroOrdine": ordini.NumeroOrdine,
		"oggetto":      ordini.Oggetto,
		"data":         primitive.NewDateTimeFromTime(time.Now()),
		"destinatario": ordini.Destinatario,
		"mittente":     ordini.Mittente,
		"prezzo":       ordini.Prezzo,
	}
	ordiniCollection := db.GetCollection(r.DB, "Ordini")
	// query
	update, err := ordiniCollection.UpdateOne(ctx, filter, bson.M{"$set": ordiniUpdate})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messaggio": err.Error()})
		return
	}
	if update.MatchedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"messaggio": "nessun ordine trovato"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "modificato con successo", "payload": ordiniUpdate})
}

// Delete Ordini
//
//	@Summary		Delete one Ordine
//	@Description	Esegui il delete di un Ordine
//	@Tags			Ordini
//	@Param			numeroOrdine	path	string	true	"Numero Ordine dell'Ordine"
//	@Produce		json
//	@Success		200
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/gest/{numeroOrdine} [delete]
func (r *Routes) DeleteOrdine(c *gin.Context) {

	numOrdine := c.Param("numeroOrdine")
	numOrdinetoInt, rerr := strconv.Atoi(numOrdine)
	if rerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "STRONZO"})
		return
	}
	filter := bson.M{"numeroOrdine": numOrdinetoInt}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	OrdiniCollection := db.GetCollection(r.DB, "Ordini")

	delete, err := OrdiniCollection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Messaggio": err.Error()})
	}
	if delete.DeletedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Nessun ordine trovato"})
	}
	c.JSON(http.StatusOK, gin.H{"Messaggio": "Ordine eliminato con successo"})
}
