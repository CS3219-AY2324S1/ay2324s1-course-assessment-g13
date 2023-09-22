package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string `json:"username"`
	Role           string `json:"role"`
	HashedPassword string `json:"-"`
}

type UserCredential struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
