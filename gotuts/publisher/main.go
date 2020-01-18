package main

import (
	"fmt"
	"time"
)

func publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		close(ch) // Broadcast to all receivers.
	}()
	return ch
}

func main() {
	wait := publish("Channels let goroutines communicate.", 850*time.Millisecond)
	fmt.Println("Waiting for news...")
	<-wait
	fmt.Println("Time to leave.")
}
