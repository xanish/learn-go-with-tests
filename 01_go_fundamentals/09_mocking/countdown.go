package _9_mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}

	_, _ = fmt.Fprint(out, finalWord)
}
