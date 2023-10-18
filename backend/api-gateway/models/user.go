package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	OauthID  int    `json:"oauth_id"`
	Provider string `json:"provider"`
	Role     string `json:"role" gorm:"default:'user'"`
}
