package _1_select

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	want := fastURL
	got, err := Racer(slowURL, fastURL)
	if err != nil {
		t.Fatalf("expected err to be nil got %v", err)
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
