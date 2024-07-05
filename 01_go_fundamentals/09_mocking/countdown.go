package _9_mocking

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from countdownStart to 1 and then prints finalWord.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	_, _ = fmt.Fprint(out, finalWord)
}

func main() {
	// Create a new ConfigurableSleeper instance with default values.
	sleeper := &ConfigurableSleeper{}

	// Call Countdown with os.Stdout as the output and the sleeper instance.
	Countdown(os.Stdout, sleeper)
}
