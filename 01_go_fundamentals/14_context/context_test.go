package _4_context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// SpyStore allows you to simulate a store and see how its used.
type SpyStore struct {
	response string
}

// Fetch returns response after a short delay.
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1) // Buffered channel to hold fetched data asynchronously

	go func() {
		var result string
		// Iterate over each character in the simulated response
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		// Send the accumulated result through the channel
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// SpyResponseWriter checks whether a response has been written. Since SpyResponseWriter
// is a mock for ResponseWriter it needs to have the methods Header, WriteHeader and Write.
type SpyResponseWriter struct {
	written bool // Flag to indicate if response has been written
}

// Header will mark written to true.
func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

// Write will mark written to true.
func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

// WriteHeader will mark written to true.
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	data := "hello, world"

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data}

		// Create a server instance using the SpyStore
		svr := Server(store)

		// Mock HTTP request and response recorder
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		// Check if the response body matches the expected data
		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// Create a context that will be cancelled after 5 milliseconds
		cancellingCtx, cancel := context.WithCancel(request.Context())
		// Cancel context after 5 milliseconds
		time.AfterFunc(5*time.Millisecond, cancel)
		// Assign the cancelling context to the request
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}
		svr.ServeHTTP(response, request)

		// Verify that no response has been written
		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
