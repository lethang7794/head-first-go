package main

import "fmt"

func main() {
	// A slice might only be able to hold values of string (or bool...)
	var _ []string
	var _ []bool

	// A map might only able to hold values of string (or bool...)
	var _ map[string]string
	var _ map[string]bool

	// A struct can hold values of many types.
	// Each value is hold in a field.

	// Declare a struct and use the struct type as the type of a variable
	var aStruct struct {
		field1 string
		field2 int
	}
	fmt.Printf("%#v\n", aStruct)

	// Declare a variable named "myStruct"
	// The "myStruct" variable can hold structs that have:
	// - a float64 "number" field
	// - a string "word" field
	// - a bool "toggle" field
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	// Print the struct value as it would appear in Go code
	fmt.Printf("%#v\n", myStruct)
	//	Each field have been set to the zero value for its type
}
