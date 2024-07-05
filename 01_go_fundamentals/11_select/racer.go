package _1_select

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// Racer performs a URL race between two URLs with a default timeout of 10 seconds.
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer performs a URL race between two URLs with a configurable timeout.
func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	// The select statement allows us to wait on multiple communication operations.
	// If multiple cases trigger at the same time select picks in random order.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// ping sends an HTTP GET request to a URL in a goroutine and returns a channel.
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()
	return ch
}
