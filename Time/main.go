package main

import (
	"fmt"
	"math"
	"time"
)

// Time is a time
type Time struct {
}

func main() {
	now := time.Now()
	fmt.Printf("starting count %v\n", now)
	j := 0
	for i := 1; i < 10000000000; i++ {
		// J increment
		j++
	}
	k := 0
	for {
		k++
		if k >= 11 {
			break
		}
	}
	fmt.Printf("ending count %v\n", time.Now())
	fmt.Printf("ending j %d\n", j)
	fmt.Printf("ending k %d\n", k)
	afterTenMiutes := now.Add(time.Minute * 10)
	dur := afterTenMiutes.Sub(now)
	if math.Abs(dur.Minutes()) < 60 {
		fmt.Println("Comparators work!")
	}
	fmt.Printf("ten seconds later %v", afterTenMiutes)
}
