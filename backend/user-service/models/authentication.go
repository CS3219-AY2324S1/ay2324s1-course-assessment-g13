package model

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Claims struct {
	User User `json:"user"`
	jwt.RegisteredClaims
}
