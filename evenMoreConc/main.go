package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func safety() {
	if r := recover(); r != nil {
		fmt.Println("recovered from error...")
	}
	wg.Done()
}

func main() {
	start := time.Now()
	for n := 2; n <= 12; n++ {
		wg.Add(1)
		go timestable(n)
		// timestable(n)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println(elapsed.Seconds())
}
func timestable(x int) {
	defer safety()
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, x, x*i)
		time.Sleep(100 * time.Millisecond)
	}
}
