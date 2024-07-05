package _3_sync

import "sync"

// Counter will increment a number. This problem could've easily been solved
// using channels as well but mutexes just make handling this case much cleaner
// compared to waiting on a channel.
type Counter struct {
	// Mutex ensures only one goroutine can increment the counter at a time.
	mu    sync.Mutex
	value int
}

// NewCounter returns a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}

// Inc will increment the count.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current count.
func (c *Counter) Value() int {
	return c.value
}
