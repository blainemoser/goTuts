package main

import "testing"

func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-2) + Fib(n-1)
}

func main() {
	for i := 0; i < 10; i++ {
		println(Fib(i))
	}
}
