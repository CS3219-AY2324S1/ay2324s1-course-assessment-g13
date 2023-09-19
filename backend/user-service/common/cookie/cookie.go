package cookie

import (
	"net/http"
	"time"
	"user-service/common/errors"
)

const rootPath = "/"
const (
	NOT_LOGIN_MESSAGE   = "Not Login"
	BAD_REQUEST_MESSAGE = "Bad Request"
)

func CreateCookie(name string, value string, expires time.Time) *http.Cookie {
	newCookie := new(http.Cookie)
	newCookie.Name = name
	newCookie.Value = value
	newCookie.Expires = expires
	newCookie.Path = rootPath
	return newCookie
}

func GetCookieValue(cookie *http.Cookie, err error) (string, errors.ServiceError) {
	if err != nil {
		if err == http.ErrNoCookie {
			return "", errors.UnauthorisedError(NOT_LOGIN_MESSAGE)
		}
		return "", errors.NewServiceError(
			BAD_REQUEST_MESSAGE,
			http.StatusBadRequest,
		)
	}
	return cookie.Value, nil
}

func SetCookieExpires(cookie *http.Cookie, err error) (*http.Cookie, errors.ServiceError) {
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, errors.UnauthorisedError(NOT_LOGIN_MESSAGE)
		}
		return nil, errors.NewServiceError(
			BAD_REQUEST_MESSAGE,
			http.StatusBadRequest,
		)
	}
	cookie.Expires = time.Now()
	return cookie, nil
}
