package main

import "fmt"

func main() {
	var myInt int

	var myIntPointer *int

	myIntPointer = &myInt

	myInt = 42

	fmt.Println(*myIntPointer)
}
