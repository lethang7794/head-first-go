package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func main() {
	fmt.Println("What is first-class functions?")
	// In a programming languages with first-class functions,
	// functions can be ASSIGNED TO VARIABLES,
	// and then CALLED from those variables.
	// And Go language support first-class functions.

	// Declare a variable with a type of "func()". This variable can hold a function.
	var myFunction func()

	// Assign the sayHi function to the variable.
	myFunction = sayHi

	// Call the function stored in the variable.
	myFunction()
}
