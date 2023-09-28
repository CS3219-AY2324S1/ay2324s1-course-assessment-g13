package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title" validate:"required"`
	Description string             `bson:"description" validate:"required"`
	Categories  []string           `bson:"categories" validate:"min=1,dive,min=5"`
	Complexity  string             `bson:"complexity" validate:"oneof=Easy Medium Hard"`
}

type EditRequest struct {
	Title       string   `bson:"title,omitempty"`
	Description string   `bson:"description,omitempty"`
	Categories  []string `bson:"categories,omitempty" validate:"omitempty,min=1,dive,min=5"`
	Complexity  string   `bson:"complexity,omitempty" validate:"omitempty,oneof=Easy Medium Hard"`
}
