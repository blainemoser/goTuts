package main

import (
	"fmt"
	"log"
)

// * This will demo error handling in Go

// Let's assume 2's are bad for some reason
func demo(c int) (int, error) {
	tally := 0
	for i := 0; i < c; i++ {
		if i == 54 {
			// return 0, errors.New("there is no such thing as 2")
			// ! this is useful for simple string errors - where one only needs one string message
			return 0, fmt.Errorf("there is no such thing as 54 - `i` is %d", i)
		}
		tally++
	}

	return tally, nil
}

type myCustomError struct {
	code    int
	detail  string
	summary string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("Error B was encountered: %s", e.summary)
}

type mySecondCustomError struct {
	secondCode int
	detail     string
	summary    string
	prediction string
}

func (e *mySecondCustomError) Error() string {
	return fmt.Sprintf("Well this is awkward")
}

// ! since errors are interfaces, we can return a struct as an error
// ! the caller can then use type-assertion to check for a specific type of error (like instance of)
func demoTwo(c int) (int, error) {
	tally := 0
	for i := 0; i < c; i++ {
		if i == 2 {
			// return 0, errors.New("there is no such thing as 2")
			return 0, &mySecondCustomError{500, "server made yucky fucky woo", "bad things", "probably some really contrived bug"}
		}
		tally++
	}

	return tally, nil
}

func main() {
	ok, err := demo(12)
	if err != nil {
		fmt.Println("Error encountered in demo function: ", err)
		return
	}
	// ! To use the data in a custom error, get the error as an instance of the custom error type via type assertion
	okTwo, errTwo := demoTwo(34)
	fmt.Println(errTwo)
	// ! here we are type-asserting the error to access its properties
	if errTwo, isMySecondCustomError := errTwo.(*mySecondCustomError); isMySecondCustomError {
		// ? A note on this syntax: if the type-assertion succeeds, then proceed in this code block given that 'myCustomError' (the first var) is of the type asserted
		// ? Therefor we can access the types properties etc.
		// This will print out a fatal error - no further execution. ? Will defer work here? Maybe use panic instead....
		log.Fatalf("This is a fatal error. Code: %d; Summary: %s; Detail: %s; Prediction: %s", errTwo.secondCode, errTwo.summary, errTwo.detail, errTwo.prediction)
		fmt.Println("this code will not be executed")
	}

	// more standard var names
	if errTwo, ok := errTwo.(*myCustomError); ok {
		fmt.Println(errTwo.code)
		fmt.Println(errTwo.detail)
	}

	fmt.Printf("Answer is %d, %d", ok, okTwo)
	// JSON decode error example:
	// if err := dec.Decode(&val); err != nil {
	// 	if serr, ok := err.(*json.SyntaxError); ok {
	// 		line, col := findLine(f, serr.Offset)
	// 		return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
	// 	}
	// 	return err
	// }
}
