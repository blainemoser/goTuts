package main

import "fmt"

func main() {

	type I interface {
		f1(name string)
		f2(name string) (error, float32)
		f3() int64
	}

	// Maps (these are dictionaries or associative arrays)
	var m = map[int]string{
		1: "my",
		2: "name",
		3: "is",
		4: "blaine",
	}

	// How to delete an element in a collective
	delete(m, 2)

	for i := range m {
		fmt.Println("i:", i)
		fmt.Println("v:", m[i])
	}
	fmt.Println(m[1])

	// Slices
	s := make([]string, 3)
	s[1] = "here"
	s[2] = "there"
	fmt.Println("slices", s)
	fmt.Println("the length of this slice is: ", len(s))
	if s[0] == "" {
		fmt.Println("yes, this is an empty string at position '0'")
	}

	// Append a value into the slice
	s = append(s, "everywhere")
	fmt.Println("the new slice looks like this: ", s)
	s = append(s, "everything")
	fmt.Println("...and now it looks like this. We have added a new element dispite our initial defining of the slice's order as 3: ", s)
	fmt.Println("now for more elements... ", "a, b and c")
	s = append(s, "a", "b", "c")
	fmt.Println("Now it looks like this: ", s)
	fmt.Println("Here is a cross-section of the slice from index 2 to index 4 (not including)", s[2:4])

	var s2 []float32
	s2 = append(s2, 1)
	s2 = append(s2, 32)
	fmt.Println("NEW slice at '0'", s2[0])
	fmt.Println("NEW slice at '1'", s2[1])

	// Arrays

	// This is an array of 12 integers
	var arr [12]int
	arr[1] = 12
	fmt.Println("array at pos. '1' is: ", arr[1])
	fmt.Println("array at pos. '11' (the last element) is: ", arr[11])

	// INDEX features
	fmt.Println("index features")

	erty := [...]int{1, 2, 3, 4, 5}
	fmt.Println(erty[2:]) // same as erty[2 : len(a)] - INCLUDES 2
	fmt.Println(erty[:3]) // same as erty[0 : 3] - INCLUDES 0
	fmt.Println(erty[:])  // same as erty[0 : len(a)] - INCLUDES THE 0 AND END POSITION + 1

	// This is a slice of integers
	intSlicer := []int{1, 2, 3, 4, 5}
	intSlicerCross := intSlicer[3:]
	fmt.Println(intSlicerCross)

	// Using the copy function - make

	// make a slice of ints that is 10 elements long and has five blank elements leading
	intSlicerCrossTwo := make([]int, 5, 10)
	copy(intSlicerCrossTwo, intSlicerCross)
	fmt.Println("used the `copy` function")
	fmt.Println(intSlicerCrossTwo)
	intSlicerCrossTwo = append(intSlicerCrossTwo, 1, 12, 23, 34)
	fmt.Println(intSlicerCrossTwo)
}
