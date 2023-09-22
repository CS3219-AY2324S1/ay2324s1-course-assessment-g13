package cookie

import (
	"api-gateway/utils/expiry"
	"net/http"
	"time"
)

const (
	FAILURE_COOKIE_NOT_FOUND = "No Cookie Found, Please Login!"
	FAILURE_GETTING_COOKIE   = "An Error Occured while Trying to Get Cookie"
	SUCCESS_COOKIE_FOUND     = "Cookie Value Found"
	SUCCESS_COOKIE_EXPIRED   = "Cookie has Expired"
)

type CookieService interface {
	CreateCookie(name string, value string, expirationTime time.Time) (cookie *http.Cookie)
	GetCookieValue(cookie *http.Cookie, err error) (cookieValue string, statusCode int, message string)
	SetCookieExpires(oldCookie *http.Cookie, err error) (newCookie *http.Cookie, statusCode int, message string)
}

type cookieService struct{}

var Service = CreateCookieService()

func (*cookieService) CreateCookie(name string, value string, expirationTime time.Time) (cookie *http.Cookie) {
	cookie = new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expirationTime
	cookie.HttpOnly = true
	return cookie
}

func (*cookieService) GetCookieValue(cookie *http.Cookie, err error) (cookieValue string, statusCode int, message string) {
	if err != nil {
		if err == http.ErrNoCookie {
			return "", http.StatusUnauthorized, FAILURE_COOKIE_NOT_FOUND
		}
		return "", http.StatusInternalServerError, FAILURE_GETTING_COOKIE
	}
	return cookie.Value, http.StatusOK, SUCCESS_COOKIE_FOUND
}

func (service *cookieService) SetCookieExpires(oldCookie *http.Cookie, err error) (newCookie *http.Cookie, statusCode int, message string) {
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, http.StatusUnauthorized, FAILURE_COOKIE_NOT_FOUND
		}
		return nil, http.StatusInternalServerError, FAILURE_GETTING_COOKIE
	}
	expireNow := expiry.ExpireNow()
	newCookie = service.CreateCookie(oldCookie.Name, oldCookie.Value, expireNow)
	return newCookie, http.StatusOK, SUCCESS_COOKIE_EXPIRED
}

func CreateCookieService() CookieService {
	return &cookieService{}
}
