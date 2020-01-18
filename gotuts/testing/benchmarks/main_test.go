package main_test

import "testing"

// Let's write a benchmark for the above function "fib"
func BenchmarkFib(b *testing.B) {
	for i := 0; i < 10; i++ {
		Fib(i)
	}
}
