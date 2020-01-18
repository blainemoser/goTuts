package main

import (
	"fmt"
	_ "greet/greeter"
)

func main() {
	fibTen := fib(10)
	fmt.Printf("Fibonacci of 10 is %d\n", fibTen)
}
