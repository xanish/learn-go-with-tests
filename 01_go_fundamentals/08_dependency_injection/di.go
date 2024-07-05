package _8_dependency_injection

import (
	"fmt"
	"io"
	"net/http"
)

// Greet sends a personalized greeting to the provided writer.
func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler is an HTTP handler function that greets the world.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
