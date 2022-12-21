package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //bson: It is a binary-encoded serialization of JSON documents
	Name    string             `json:"name,omitempty"`
	Studied bool               `json:"studied,omitempty"`
}
