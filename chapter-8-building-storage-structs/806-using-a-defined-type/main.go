package main

import "fmt"

// Define a type name Subcriber
type Subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	var subscriber1 Subscriber
	subscriber1.name = "Harry"

	var subscriber2 Subscriber
	subscriber2.name = "Hermione"

	fmt.Printf("%#v\n", subscriber1)
	fmt.Printf("%#v\n", subscriber2)
}
