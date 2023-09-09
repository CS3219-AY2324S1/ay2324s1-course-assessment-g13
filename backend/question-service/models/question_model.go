package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	Categories  []string           `json:"categories" bson:"categories" validate:"min=1,dive,min=5"`
	Complexity  string             `json:"complexity" bson:"complexity" validate:"oneof=Easy Medium Hard"`
}

type EditRequest struct {
	Title       string   `json:"title,omitempty" bson:"title,omitempty"`
	Description string   `json:"description,omitempty" bson:"description,omitempty"`
	Categories  []string `json:"categories,omitempty" bson:"categories,omitempty" validate:"omitempty,min=1,dive,min=5"`
	Complexity  string   `json:"complexity,omitempty" bson:"complexity,omitempty" validate:"omitempty,oneof=Easy Medium Hard"`
}
