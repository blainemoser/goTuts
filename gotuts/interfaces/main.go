package main

import (
	"fmt"
	"math"
)

/*
	An interface is a type that holds a set of methods that can be called for a variety of types
*/

// Animal ... This interface 'Animal' has the method speak. Thus, an 'animal' is that which can speak
type Animal interface {
	Speak() string
}

// Dog ...
type Dog struct {
}

// Speak ...
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat ...
type Cat struct {
}

// Speak ... note that this is a pointer to Cat
func (c *Cat) Speak() string {
	return "Meow!"
}

// Llama ...
type Llama struct {
}

// Speak ...
func (l Llama) Speak() string {
	return "?????"
}

// JavaProgrammer ...
type JavaProgrammer struct {
}

// Speak ...
func (j JavaProgrammer) Speak() string {
	return "Design Patterns!"
}

// DoSomething ...
func DoSomething(v interface{}) {
	// This function will accept any parameter - because all types implement 0 methods
	// Use type checking to see whether the input is iterable
	// if v, ok := v.([]string); ok {
	// 	for _, value := range v {
	// 		fmt.Println(value)
	// 	}
	// 	return
	// }
	// Or use switch with a default handler
	fmt.Println(math.Pow(2, 3))
	switch x := v.(type) {
	case string:
		fmt.Println(x)
		return
	case []string:
		for _, value := range x {
			fmt.Println(value)
		}
		return
	case int:
		fmt.Println(x)
	default:
		// Pass silently
		return
	}
	// fmt.Println(v)
	// return
}

func main() {
	animals := []Animal{new(Dog), new(Cat), Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	DoSomething(12)       // Call DoSomething on an integer
	DoSomething("Blaine") // Call DoSomething on a string
	DoSomething([]string{"List Elem 1", "List Elem 2", "List Elem 3"})
}
