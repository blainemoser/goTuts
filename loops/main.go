package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		return fib(n-2) + fib(n-1)
	}
}

func main() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}

	// Here is some alternative syntax
	j := 0
	for j <= 10 {
		fmt.Println("alternate: ", j)
		j++
	}

	/*
		Note that in Go there is no while loop specifier.
		One can construct an infinite loop as:
		for {
			#code....
		}

		Use a break statement to complete the loop
	*/
	x := 0
	for {
		fmt.Println("Do stuff... ", x)
		x++
		// Add a break condition
		if x > 9 {
			break
		}
	}

	// Use multiple vars in a loop
	b := 0
	for y := 5; b < 25; y += 3 {
		fmt.Println("multi-var", y)
		b++
	}

	// Fibbin' it up
	for d := 1; d < 11; d++ {
		fmt.Printf("Fibonacci of %d is %d \n", d, fib(d))
	}
}
