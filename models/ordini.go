package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ordini struct {
	NumeroOrdine int                `json:"numeroOrdine" bson:"numeroOrdine", omitempty`
	Oggetto      string             `json:"oggetto" bson:"oggetto"`
	Data         primitive.DateTime `json:"data" bson:"data", omitempty`
	Destinatario string             `json:"destinatario" bson:"destinatario"`
	Mittente     string             `json:"mittente" bson:"mittente"`
	Prezzo       float64            `json:"prezzo" bson:"prezzo"`
}

type Counter struct {
	Id  string `json:"_id" bson:"_id"`
	Seq int    `json:"seq" bson:"seq", omitempty`
}
