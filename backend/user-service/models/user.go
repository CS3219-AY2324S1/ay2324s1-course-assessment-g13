package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string
	HashedPassword string
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
