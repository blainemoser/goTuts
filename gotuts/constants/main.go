package main

import "fmt"

const (
	monday = 1 << iota
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

func main() {
	fmt.Println(sunday)
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(wednesday)
}
