package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID `bson:"_id"`
	Author      *string            `json:"author" validate:"required"`
	Title       *string            `json:"title" validate:"required"`
	Description *string            `json:"description"	validate:"required"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
}
