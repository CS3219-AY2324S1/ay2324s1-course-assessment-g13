package auth

import "time"

func GetExpirationTime() time.Time {
	return time.Now().Add(5 * time.Minute)
}
