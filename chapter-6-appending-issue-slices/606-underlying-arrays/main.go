package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}

	ac := array[:3] // [a b c]
	de := array[3:] // [d e]

	_ = array[:3]   // [a b c]
	ce := array[2:] // [c d e]

	fmt.Println(ac, de, ce)
}
