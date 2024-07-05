package _0_concurrency

import "net/http"

// CheckWebsite returns true if the URL returns a 200 status code, false otherwise.
func CheckWebsite(url string) bool {
	// Ideally we should not be using http.(Get / Head / Post) as these
	// use the DefaultClient which does not have configured timeouts.
	// But its fine here since this is a test program.
	response, err := http.Head(url)
	if err != nil {
		return false
	}

	return response.StatusCode == http.StatusOK
}
