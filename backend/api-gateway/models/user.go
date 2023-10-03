package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   int    `json:"id"`
	Provider string `json:"provider"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Picture  string `json:"picture"`
	Role     string `json:"role" gorm:"default:'user'"`
}
