package auth

import (
	"os"
	"time"
	"user-service/common/constants"
	"user-service/common/errors"
	model "user-service/models"

	"github.com/golang-jwt/jwt/v5"
)

type TokenServiceI interface {
	Generate(user *model.User, expirationTime time.Time) (string, errors.ServiceError)
	Validate(tokenString string) (*model.Claims, errors.ServiceError)
}

type tokenService struct {
	jwtSecretKey []byte
}

var TokenService = CreateTokenService()

func (service *tokenService) Generate(user *model.User, expirationTime time.Time) (string, errors.ServiceError) {
	claims := &model.Claims{
		User: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(service.jwtSecretKey)
	if err != nil {
		return "", errors.InternalServerError()
	}
	return tokenString, nil
}

func (service *tokenService) Validate(tokenString string) (*model.Claims, errors.ServiceError) {
	claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return service.jwtSecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.UnauthorisedError("Invalid Signature")
		}
		return nil, errors.InvalidTokenError()
	}

	if !token.Valid {
		return nil, errors.UnauthorisedError("Token Expires")
	}

	return claims, nil
}

func CreateTokenService() TokenServiceI {
	secretKey := os.Getenv(constants.JWT_SECRET_KEY_ENV)
	return &tokenService{
		jwtSecretKey: []byte(secretKey),
	}
}
