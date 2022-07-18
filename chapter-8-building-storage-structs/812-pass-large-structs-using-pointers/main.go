package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) { // Take a pointer.
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) *subscriber { // Return a pointer.
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s // Return a point to the struct instead of the struct itself.
}

func main() {
	subscriber1 := defaultSubscriber("Harry Potter")
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Hermione Granger")
	subscriber2.rate = 4.99
	printInfo(subscriber2)
}
