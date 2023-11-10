package model

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	AuthUserID        uint      `json:"auth_user_id"`
	Username          string    `json:"username"`
	PhotoUrl          string    `json:"photo_url"`
	PreferredLanguage string    `json:"preferred_language"`
	Histories         []History `gorm:"foreignKey:UserID"`
}

type CreateUser struct {
	AuthUserID        uint   `json:"auth_user_id" validate:"required"`
	Username          string `json:"username" validate:"required"`
	PhotoUrl          string `json:"photo_url"`
	PreferredLanguage string `json:"preferred_language"`
}

type UpdateUser struct {
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
	UserID uint `json:"user_id" gorm:"index;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreateHistory struct {
	RoomId string `json:"room_id" validate:"required"`
	QuestionId string `json:"question_id" validate:"required"`
	Title string `json:"title" validate:"required"`
	Solution string `json:"solution" validate:"required"`
	Username string `json:"username" validate:"required"`
}