package cookie

import (
	"net/http"
	"time"
	"user-service/common/errors"
)

func CreateCookie(name string, value string, expires time.Time) *http.Cookie {
	newCookie := new(http.Cookie)
	newCookie.Name = name
	newCookie.Value = value
	newCookie.Expires = expires
	return newCookie
}

func GetCookieValue(cookie *http.Cookie, err error) (string, errors.ServiceError) {
	if err != nil {
		if err == http.ErrNoCookie {
			return "", errors.UnauthorisedError("Not Login")
		}
		return "", errors.NewServiceError(
			"Bad Request",
			http.StatusBadRequest,
		)
	}
	return cookie.Value, nil
}

func SetCookieExpires(cookie *http.Cookie, err error) (*http.Cookie, errors.ServiceError) {
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, errors.UnauthorisedError("Not Login")
		}
		return nil, errors.NewServiceError(
			"Bad Request",
			http.StatusBadRequest,
		)
	}
	cookie.Expires = time.Now()
	return cookie, nil
}
