package main

import "fmt"

// Use 'defer' to specify what must be run last in a function.
// Defer is run on a Last-In-First-Out basis - so defer will be run in descending order here
func foo() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	foo()
}
