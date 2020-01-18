package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// cleanup function
func safetyNet() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from error f ", r)
	}
	wg.Done()
}

// Create a function that takes a channel
func foo(c chan int, someValue int) {
	defer safetyNet()
	c <- someValue * 5
}

func main() {
	// This is how to make a channel
	fooVal := make(chan int, 10) // This channel has a buffer of 10 - check the second parameter
	// Return the first 10 digits
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i)
	}

	// Make sure that there is a wait here
	wg.Wait()

	// close the channel after use
	close(fooVal)

	// Use 'range' to iterate through the channel values
	for item := range fooVal {
		fmt.Println(item)
	}
}
