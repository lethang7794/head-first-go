package main

import "fmt"

func calmDown() {
	recover()
}

func freakOut2() {
	defer calmDown()               // Defer a call to the function that recovers.
	panic("oh no")                 // If the program panics AFTER that, the deferred function call will recover!
	fmt.Println("I won't be run!") // The code after the panic will never be run
}

func main() {
	freakOut2()
	fmt.Println("Exiting normally!") // But this code runs after freakOut returns.
}
