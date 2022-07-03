package main

import "fmt"

// Declare that the function return a float64 pointer
func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat // Pointer to a variable that's local to a function
}

// Use a pointer type for function param
func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func main() {
	var myFloatPointer *float64 = createPointer()
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	var myBool bool = true
	printPointer(&myBool)
}
