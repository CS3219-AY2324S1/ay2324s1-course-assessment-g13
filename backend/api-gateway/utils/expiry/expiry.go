package expiry

import (
	"time"
)

func ExpireNow() time.Time {
	return time.Now()
}

func ExpireIn5Minutes() time.Time {
	return time.Now().Add(5 * time.Minute)
}

func ExpireIn24Hours() time.Time {
	return time.Now().Add(24 * time.Hour)
}
