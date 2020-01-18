package main

import "fmt"

func augment(x *int) {
	*x = *x + 6
	return
}

func main() {
	x := 5
	a := &x
	// j := &x
	// This demonstrates that the object in memory is being altered
	*a = (*a) * (*a)
	fmt.Println(a)  // Prints the location in memory address
	fmt.Println(*a) // Prints the value of a - which is x also
	fmt.Println(x)
	augment(&x)
	fmt.Println("Here is x after the augment function:", x)
}
