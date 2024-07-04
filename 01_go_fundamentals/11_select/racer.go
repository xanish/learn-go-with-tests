package _1_select

import (
	"net/http"
	"time"
)

func Racer(a, b string) (string, error) {
	startA := time.Now()
	_, err := http.Get(a)
	if err != nil {
		return "", err
	}
	aDuration := time.Since(startA)

	startB := time.Now()
	_, err = http.Get(b)
	if err != nil {
		return "", err
	}
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a, nil
	}

	return b, nil
}
