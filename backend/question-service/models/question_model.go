package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" validate:"required"`
	Categories  []string           `json:"categories" bson:"categories" validate:"min=1,dive,min=5"`
	Description string             `json:"description" bson:"description" validate:"required"`
	Complexity  string             `json:"complexity" bson:"complexity" validate:"oneof=Easy Medium Hard"`
}
