package expiry

import (
	"time"
)

func ExpireNow() time.Time {
	return time.Now()
}

func ExpireIn1Minute() time.Time {
	return time.Now().Add(time.Minute)
}

func ExpireIn24Hours() time.Time {
	return time.Now().Add(24 * time.Hour)
}
