package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	OauthID       int    `json:"oauth_id"`
	OauthProvider string `json:"oauth_provider"`
	Role          string `json:"role" gorm:"default:'user'"`
}

type LoginRequest struct {
	OauthID       int    `json:"oauth_id" validate:"required"`
	OauthProvider string `json:"oauth_provider" validate:"required"`
}

type CreateUser struct {
	OauthID           int    `json:"oauth_id" validate:"required"`
	OauthProvider     string `json:"oauth_provider" validate:"required"`
	Username          string `json:"username" validate:"required"`
	PhotoUrl          string `json:"photo_url"`
	PreferredLanguage string `json:"preferred_language"`
}

type UserServiceCreateUserRequestBody struct {
	AuthUserID        uint   `json:"auth_user_id"`
	Username          string `json:"username"`
	PhotoUrl          string `json:"photo_url"`
	PreferredLanguage string `json:"preferred_language"`
}

type UpdateUser struct {
	Username          string `json:"username"`
	PhotoUrl          string `json:"photo_url"`
	PreferredLanguage string `json:"preferred_language"`
}

type UserServiceUser struct {
	gorm.Model
	AuthUserID        uint   `json:"auth_user_id"`
	Username          string `json:"username"`
	PhotoUrl          string `json:"photo_url"`
	PreferredLanguage string `json:"preferred_language"`
}

type UpdateUserResponse struct {
	Message string `json:"message"`
	User	UserServiceUser   `json:"user"`
}
