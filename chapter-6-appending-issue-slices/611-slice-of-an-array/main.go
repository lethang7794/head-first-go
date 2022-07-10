package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[1:3]             // [b c]
	slice = append(slice, "x")      // [b c x]
	slice = append(slice, "y", "z") // [b c x y z]

	for _, letter := range slice {
		fmt.Println(letter)
	}
}
