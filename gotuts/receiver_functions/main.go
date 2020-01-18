package main

import "fmt"

type person struct {
	name string
	age  int
}

type dog struct {
	name string
	age  int
}

type animal interface {
	details()
}

// Receiver function for person
func (p person) details() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

// Receiver function for person
func (d dog) details() {
	fmt.Printf("%s is %d years old", d.name, d.age)
}

func main() {
	blaine := person{
		name: "Blaine",
		age:  30,
	}
	blaine.details()
}
