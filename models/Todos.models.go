package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty" form:"title"`
	Description string             `bson:"description,omitempty" form:"description"`
	UserId      string             `bson:"userId,omitempty" form:"userid"`
}