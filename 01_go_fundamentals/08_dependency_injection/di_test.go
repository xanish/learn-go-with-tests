package _8_dependency_injection

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreet(t *testing.T) {
	// The Buffer type from the bytes package implements the Writer interface,
	// because it has the method Write(p []byte) (n int, err error).
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMyGreeterHandler(t *testing.T) {
	// Create a new HTTP request using http.NewRequest.
	request, _ := http.NewRequest(http.MethodGet, "/", nil)

	// Create a new HTTP recorder to record the response.
	response := httptest.NewRecorder()

	// Create an HTTP handler function from MyGreeterHandler.
	handler := http.HandlerFunc(MyGreeterHandler)

	// Call the HTTP handler function with the response recorder and request.
	handler.ServeHTTP(response, request)

	// Check the response captured in the recorder.
	got := response.Body.String()
	want := "Hello, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
