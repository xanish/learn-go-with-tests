package _9_mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

// SpyTime is a spy implementation for tracking sleep durations.
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep records the duration slept.
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// SpyCountdownOperations is a spy implementation for tracking method calls.
type SpyCountdownOperations struct {
	Calls []string
}

// Sleep records the sleep operation call.
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write records the write operation call.
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}

		// Since SpyCountdownOperations has both Sleep and Write defined
		// on it, we can pass it as both writer and sleeper to mock their
		// functionality and then check if the order of operations
		// happened correctly.
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}

	// Pass the sleep duration and a spy sleeper function from spyTime.Sleep
	// which mocks functionality of Sleep without actually sleeping.
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
