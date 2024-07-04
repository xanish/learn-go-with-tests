package _4_context

import (
	"fmt"
	"net/http"
)

// Store fetches data.
type Store interface {
	Fetch() string
}

// Server returns a handler for calling Store.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, store.Fetch())
	}
}
