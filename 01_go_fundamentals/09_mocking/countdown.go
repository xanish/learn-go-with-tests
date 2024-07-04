package _9_mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper is an implementation of Sleeper with a defined delay.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep will pause execution for the defined Duration.
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	_, _ = fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{}
	Countdown(os.Stdout, sleeper)
}
