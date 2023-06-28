package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
}
