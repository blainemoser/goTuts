package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	stringEx := "Blaine Moser"

	// Let's demo some strings functions
	fmt.Println(strings.Contains(stringEx, "Bl"))
	fmt.Println(strings.Index(stringEx, "Bl"))
	fmt.Println(strings.Count(stringEx, "Bl"))
	fmt.Println(strings.Replace(stringEx, "Bl", "ML", 25))

	// Let's demo some strings/sort functions
	csvStrings := "1,2,3,4,5"

	// Creates a slice of elements, split by the delimiter string
	fmt.Println(strings.Split(csvStrings, ","))

	listOfLetters := []string{"d", "e", "r", "o", "u", "b", "a"}
	// Sort the list of letters
	sort.Strings(listOfLetters)
	fmt.Println(listOfLetters)

	// Create a string from an array - like implode
	stringsFromArr := strings.Join(listOfLetters, ", ")
	fmt.Println(stringsFromArr)

	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75},
		{"Bob", 75},
		{"Alice", 75},
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}

	// Sort by name, preserving original order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	// Sort by age preserving name order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age,name:", people)
}
