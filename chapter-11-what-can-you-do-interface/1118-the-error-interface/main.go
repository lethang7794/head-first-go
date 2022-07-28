package main

import "fmt"

// Define a type with an underlying of "string"
type ComedyError string

// Satisfy the error interface.
func (c ComedyError) Error() string {
	return string(c) // The Error method needs to return a string, so do a type conversion.
}

func main() {
	var err error

	// ComedyError satisfies the error interface,
	// so we can assign a ComedyError to the variable.
	err = ComedyError("What's a programmer's favorite beer? Logger!")

	fmt.Println(err)
}
