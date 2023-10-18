package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	OauthID       int    `json:"oauth_id"`
	OauthProvider string `json:"oauth_provider"`
	Role          string `json:"role" gorm:"default:'user'"`
}

type UserRequestPayload struct {
	OauthID       int    `json:"oauth_id" validate:"required"`
	OauthProvider string `json:"oauth_provider" validate:"required"`
}
