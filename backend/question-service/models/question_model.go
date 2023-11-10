package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	Categories  []string           `json:"categories" bson:"categories" validate:"required"`
	Complexity  string             `json:"complexity" bson:"complexity" validate:"oneof=Easy Medium Hard"`
}

type Category struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Category string             `json:"category" bson:"category"`
}

type LeetCodeProblem struct {
	Id          string             `json:"id"`
	Title       string             `json:"title"` 
	Description string             `json:"description"`
	Categories  []string           `json:"categories"`
	Complexity  string             `json:"complexity"`
}

type APIResponse struct {
	Total    int               `json:"total"`
	Problems []LeetCodeProblem `json:"problems"`
}
