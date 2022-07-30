package main

import "fmt"

func freakOut() {
	panic("oh no") // The panic stops the rest of the function from running
	recover()      // ...so this will never be run!
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
