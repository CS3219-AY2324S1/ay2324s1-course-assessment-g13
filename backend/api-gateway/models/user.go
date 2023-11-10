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

type History struct {
	gorm.Model
	RoomId string `json:"room_id"`
	QuestionId string `json:"question_id"`
	Title string `json:"title"`
	Solution string `json:"solution"`
	Language string `json:"language"`
	UserID uint `json:"user_id"`
}

type HistoryResponse  struct {
	Message string `json:"message"`
	Histories []History `json:"histories"`
}
