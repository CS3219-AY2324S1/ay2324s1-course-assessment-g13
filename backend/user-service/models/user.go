package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID         uint
	Username       string
	HashedPassword string
	PhotoUrl       string
}

type CreateUserRequest struct {
	UserID   uint   `json:"user_id" validate:"required"`
	Username string `json:"username" validate:"required"`
	PhotoURL string `json:"photo_url" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserInfo struct {
	Username string `json:"username"`
	PhotoUrl string `json:"photoUrl"`
}

type UpdateUserPassword struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

type LoginResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	PhotoUrl string `json:"photoUrl"`
}
