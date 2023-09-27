package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category string

const (
	Algorithms         Category = "Algorithms"
	DataStructures     Category = "Data Structures"
	BrainTeaser        Category = "Brain Teaser"
	Strings            Category = "Strings"
	BitManipulation    Category = "Bit Manipulation"
	DynamicProgramming Category = "Dynamic Programming"
)

type Question struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	Categories  []Category         `json:"categories" bson:"categories" validate:"required"`
	Complexity  string             `json:"complexity" bson:"complexity" validate:"oneof=Easy Medium Hard"`
}

type EditRequest struct {
	Title       string     `bson:"title,omitempty" validate:"alpha"`
	Description string     `bson:"description,omitempty"`
	Categories  []Category `bson:"categories,omitempty" validate:"required"`
	Complexity  string     `bson:"complexity,omitempty" validate:"omitempty,oneof=Easy Medium Hard"`
}
