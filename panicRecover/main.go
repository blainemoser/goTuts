package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a weight group instance to operate in this namespace
var wg sync.WaitGroup

// This function is responsible for Panic recovery
func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in `cleanup`: ", r)
	}
	// Place the weight group decrementor here
	wg.Done()
}

func say(s string) {
	// define the ending of the routine using the wg (weight group)
	// Use "defer" to make sure the wg is set to Done even if there is a
	// problem with the rest of the function's execution
	defer cleanup()
	// Perform *task* 3 times
	for i := 0; i < 3; i++ {
		// Use a sleeper to demonstrate the concept
		if i == 2 {
			panic("Oh yuck! A 2!")
		}
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}

}

func main() {

	// Use go routines to execute the functions
	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("there")

	// This ensures that all the goroutines are completed before the main func finishes.
	wg.Wait()

	// Can start another set of goroutines after waiting for the first set to complete
	wg.Add(1)
	go say("hi")
	wg.Wait()
}
