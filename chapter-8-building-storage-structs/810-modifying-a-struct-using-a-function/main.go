package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func applyDiscount(s *subscriber) { // Take a pointer to a struct, not the struct itself
	s.rate = 4.99 // Update the struct field
	// We don't need to use the * operator to get the value at the pointer.
	// The dot notation to access fields works on struct pointer as well as the struct themselves.
}

func main() {
	var s subscriber
	applyDiscount(&s) // Pass a pointer, not a struct

	fmt.Println(s.rate)
}
