package _8_dependency_injection

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMyGreeterHandler(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(MyGreeterHandler)

	handler(response, request)

	got := response.Body.String()
	want := "Hello, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
