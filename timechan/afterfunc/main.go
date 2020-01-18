package main

import (
	"log"
	"time"
)

func main() {
	timer := time.AfterFunc(time.Second, func() {
		log.Println("Func has been running for more than a second...")
	})
	defer timer.Stop()
	time.Sleep(5 * time.Second)
}
