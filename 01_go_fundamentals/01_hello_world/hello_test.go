package _1_hello_world

import "testing"

func TestHello(t *testing.T) {
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello in English", func(t *testing.T) {
		got := Hello("Chris", "english")
		want := "Hello, Chris!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Chris", "spanish")
		want := "Hola, Chris!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Chris", "french")
		want := "Bonjour, Chris!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
