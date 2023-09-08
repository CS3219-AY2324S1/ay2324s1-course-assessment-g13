package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	HashedPassword string
	// Salt string `json:"salt"`
}

type CreateUserRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
