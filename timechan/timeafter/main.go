package main

import (
	"fmt"
	"time"
)

// This is NOT garbage-collected after the stop
func demoTimeAfter(AFP chan string) {
	select {
	case news := <-AFP:
		fmt.Println(news)
	case <-time.After(10 * time.Second):
		fmt.Println("No news in 10 Seconds.")
	}
}

// This is garbage-collected after the stop
func demoNewTimer(AFZ chan string) {
	for alive := true; alive; {
		timer := time.NewTimer(10 * time.Second)
		select {
		case news := <-AFZ:
			timer.Stop()
			fmt.Println(news)
		case <-timer.C:
			alive = false
			fmt.Println("No news in 10 Seconds. Service aborting.")
		}
	}
}

func main() {
	AFP := make(chan string)
	demoTimeAfter(AFP)

	AFZ := make(chan string)
	demoNewTimer(AFZ)
}
