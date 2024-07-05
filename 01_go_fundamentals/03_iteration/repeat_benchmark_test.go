package _3_iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	// When the benchmark code is executed, it runs b.N times and measures how long it takes.
	// By default, Benchmarks are run sequentially. We can use the RunParallel method to run them concurrently.
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
