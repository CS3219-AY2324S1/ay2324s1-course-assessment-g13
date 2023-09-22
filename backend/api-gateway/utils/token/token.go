package token

import (
	"api-gateway/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	ENV_JWT_SECRET_KEY = "JWT_SECRET_KEY"

	INVALID_TOKEN_SIGNATURE    = "Invalid Signature!"
	FAILURE_TOKEN_UNSIGN       = "Unable to Sign Token!"
	FAILURE_PARSING_JWT_CLAIMS = "An Error Occured when Parsing JWT Claims!"
	FAILURE_TOKEN_EXPIRED      = "Token Expired!"
	SUCCESS_TOKEN_GENERATED    = "Token Generated Successfully"
	SUCCESS_TOKEN_VALIDATED    = "Token is Valid"
)

type TokenService interface {
	Generate(user *models.User, expirationTime time.Time) (token string, statusCode int, message string)
	Validate(tokenString string) (tokenClaims *models.Claims, statusCode int, message string)
}

type tokenService struct {
	secretKey []byte
}

var Service = CreateTokenService()

func (service *tokenService) Generate(user *models.User, expirationTime time.Time) (string, int, string) {
	tokenClaims := &models.Claims{
		User: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	tokenClaimsWithHeader := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenSigned, err := tokenClaimsWithHeader.SignedString(service.secretKey)
	if err != nil {
		return "", http.StatusInternalServerError, FAILURE_TOKEN_UNSIGN
	}
	return tokenSigned, http.StatusOK, SUCCESS_TOKEN_GENERATED
}

func (service *tokenService) Validate(tokenString string) (*models.Claims, int, string) {
	tokenClaims := new(models.Claims)
	token, err := jwt.ParseWithClaims(tokenString, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return service.secretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, http.StatusUnauthorized, INVALID_TOKEN_SIGNATURE
		}
		return nil, http.StatusInternalServerError, FAILURE_PARSING_JWT_CLAIMS
	}
	if !token.Valid {
		return nil, http.StatusUnauthorized, FAILURE_TOKEN_EXPIRED
	}
	return tokenClaims, http.StatusOK, SUCCESS_TOKEN_VALIDATED
}

func CreateTokenService() TokenService {
	secret := os.Getenv(ENV_JWT_SECRET_KEY)
	return &tokenService{
		secretKey: []byte(secret),
	}
}
