package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserAuthorization struct {
	Role string `json:"role"`
}

type Claims struct {
	User User `json:"user"`
	jwt.RegisteredClaims
}

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
