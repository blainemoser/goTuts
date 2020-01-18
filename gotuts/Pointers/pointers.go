package main

import "fmt"

func main() {
	x := 5
	a := &x
	// This demonstrates that the object in memory is being altered
	*a = (*a) * (*a)
	fmt.Println(a) // Prints the location in memory address
	fmt.Println(x)
}
