package auth

import "time"

func GetExpirationTime() time.Time {
	return time.Now().Add(time.Hour)
}
