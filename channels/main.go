package main

import (
	"fmt"
	"time"
)

// Create a function that takes a channel
func foo(c chan int, someValue int) {
	time.Sleep(time.Millisecond * 3000)
	c <- someValue * 5
}

func main() {
	// This is how to make a channel
	fooVal := make(chan int)
	go foo(fooVal, 4)
	go foo(fooVal, 3)

	// Note that sending and receiving on a channel automatically blocks the output until receipt of the channel output
	v1 := <-fooVal
	v2 := <-fooVal
	fmt.Println(v1, v2)
	fmt.Println("done")
}
