package main

import "fmt"

func main() {
	// Declare a slice of string
	var notes []string

	// The slice hasn't created yet. Now, it has the zero value (which is nil for slice).
	fmt.Printf("%#v\n", notes)

	// Create a slice an assign it to a var
	notes = make([]string, 7)
	fmt.Printf("%#v\n", notes)

	// Assign slice's elements
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	// Retrieves slice's elements
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	// Get a slice's length with len
	fmt.Println(len(notes))

	// len also accept a nil slice
	var nilSlice []string
	fmt.Println(len(nilSlice))

	// for and for...range work just the same with slices as they do with arrays
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
	for i, v := range notes {
		fmt.Println(i, v)
	}
}
