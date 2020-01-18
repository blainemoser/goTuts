package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func doSomething() {
	fmt.Println("doing something...")
	waitGroup.Done() // remove a wait item from the queue
}

func main() {
	fmt.Println("I'm in...")
	waitGroup.Add(1) // specify that we need to wait for one goRoutine to terminate
	go doSomething()
	fmt.Println("I'm out.")
	waitGroup.Wait()
}
