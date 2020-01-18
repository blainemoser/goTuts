package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Create a sample file
	file, err := os.Create("sample_file.txt")
	if err != nil {
		// Log any fatal error
		log.Fatal(err)
	}

	// Write some text to the file
	file.WriteString("Lorum ipsem dolor sit amet")

	// Close the file
	file.Close()

	// Try and re-open the file
	stream, err := ioutil.ReadFile("sample_file.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Convert the data stream into a string
	readString := string(stream)

	fmt.Println(readString)
}
