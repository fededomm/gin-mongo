package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ordini struct {
	NumeroOrdine int                `json:"numeroOrdine" bson:"numeroOrdine", omitempty`
	Oggetto      string             `json:"oggetto" bson:"oggetto", omitempty`
	Data         primitive.DateTime `json:"data" bson:"data", omitempty`
	Destinatario string             `json:"destinatario" bson:"destinatario", omitempty`
	Mittente     string             `json:"mittente" bson:"mittente", omitempty`
	Prezzo       float64            `json:"prezzo" bson:"prezzo", omitemty`
}
