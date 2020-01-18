package main

import "fmt"

func main() {
	var i int
	y := "Bl"
	switch y {
	case "Bl":
		i = 43
		break
	case "Nl":
		i = 27
	}
	fmt.Println(i)
}
