package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}

	ac := array[:3] // [a b c]
	ce := array[2:] // [c d e]

	ae := array[:]
	bd := array[1:4]

	fmt.Println(ac, ce, ae, bd)

	// Assign a new value to a slice element will change the corresponding element in the underlying array
	ac[2] = "C"

	// The change will be visible to ALL the slices.
	fmt.Println(ac, ce, ae, bd)

	// => It's better to create slices using make or slice literal
}
