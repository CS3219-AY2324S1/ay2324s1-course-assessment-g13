package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	AuthUserID        uint   `json:"auth_user_id"`
	Username          string `json:"username"`
	PhotoUrl          string `json:"photo_url"`
	PreferredLanguage string `json:"preferred_language"`
}

type CreateUserPayload struct {
	AuthUserID        uint   `json:"auth_user_id" validate:"required"`
	Username          string `json:"username" validate:"required"`
	PhotoUrl          string `json:"photo_url"`
	PreferredLanguage string `json:"preferred_language"`
}

type UpdateUserPayload struct {
	Username          string `json:"username" validate:"atleastonefield"`
	PhotoUrl          string `json:"photo_url" validate:"atleastonefield"`
	PreferredLanguage string `json:"preferred_language" validate:"atleastonefield"`
}
