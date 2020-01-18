package main

import (
	"fmt"
	"strconv"
)

func main() {
	// randInt := 20
	// randFloat := 34.23

	randStringNumber := "342"
	randStringNumberTwo := "234.43"

	// Create and integer from the string that can be converted to a float
	newInt, _ := strconv.ParseInt(randStringNumber, 0, 64)
	fmt.Println(newInt)

	newFloat, _ := strconv.ParseFloat(randStringNumberTwo, 64)
	fmt.Println(newFloat)
}
