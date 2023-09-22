package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username               string `json:"username,omitempty"`
	Role                   string `json:"role" gorm:"default:'user'"`
	HashedPassword         string `json:"-"`
	OAuthProvider          string `json:"oauthProvider,omitempty"`
	OAuthUserID            int    `json:"-"`
	OAuthUsername          string `json:"oauthUsername,omitempty"`
	OAuthEmail             string `json:"-"`
	OAuthProfilePictureURL string `json:"-"`
	OAuthProfileURL        string `json:"-"`
}

type UserCredential struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
