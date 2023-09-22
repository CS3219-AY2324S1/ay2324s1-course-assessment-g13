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

func Add5MoreSeconds(t time.Time) time.Time {
	return t.Add(5 * time.Second)
}
