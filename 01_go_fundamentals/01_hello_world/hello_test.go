package _1_hello_world

import "testing"

func TestHello(t *testing.T) {
	got := Hello("World")
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
