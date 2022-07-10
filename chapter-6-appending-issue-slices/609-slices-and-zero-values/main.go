package main

import "fmt"

func main() {
	// Create slices without assigning values to their elements.
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)

	// Access a slice element that no value has been assigned to.
	fmt.Println(floatSlice[0], boolSlice[9])

	// Declare slice variables without create slices.
	var intSlice []int
	var stringSlice []string

	// Slice variables that no slice has been assigned to have a zero value of nil.
	fmt.Printf("%#v, %#v\n\n", intSlice, stringSlice)

	// Nil slice is treated as an empty slice
	var nilSlice []int = nil
	var emptySlice []int = []int{}
	fmt.Printf("%#v, %#v\n", nilSlice, emptySlice)
	fmt.Println(nilSlice, emptySlice)
	fmt.Println(len(nilSlice), len(emptySlice))

	// _ = nilSlice[0] // panic: runtime error: index out of range [0] with length 0
	// _ = emptySlice[0] // panic: runtime error: index out of range [0] with length 0

	// Add elements of an empty slice to the empty slice
	emptySlice = append(emptySlice, emptySlice...)

	// Add elements of a nul slice to the nil slice
	nilSlice = append(nilSlice, nilSlice...)
}
