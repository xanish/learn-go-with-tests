package _8_dependency_injection

import (
	"fmt"
	"io"
)

import (
	"net/http"
)

// Greet sends a personalised greeting to writer.
func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
