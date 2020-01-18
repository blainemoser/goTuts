package main

import "fmt"

func main() {
	i := 0
	go func() {
		i++ // write
	}()
	fmt.Println(i) // concurrent read
}

/**
RUN THE COMMAND WITH --race

*/
