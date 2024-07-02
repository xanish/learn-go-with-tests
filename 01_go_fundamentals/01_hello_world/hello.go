package _1_hello_world

import (
	"fmt"
	"strings"
)

const (
	spanish      = "spanish"
	french       = "french"
	englishHello = "Hello"
	spanishHello = "Hola"
	frenchHello  = "Bonjour"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	var greeting string
	switch strings.ToLower(language) {
	case spanish:
		greeting = spanishHello
	case french:
		greeting = frenchHello
	default:
		greeting = englishHello
	}

	return fmt.Sprintf("%s, %s!", greeting, name)
}
