package _1_select

import (
	"math"
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration, _ := measureResponseTime(a)
	bDuration, _ := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) (time.Duration, error) {
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		return math.MaxInt, err
	}
	return time.Since(start), nil
}
