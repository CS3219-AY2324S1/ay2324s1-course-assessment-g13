package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string
	Role           string `json:"role" gorm:"default:'basic'"`
	HashedPassword string
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
