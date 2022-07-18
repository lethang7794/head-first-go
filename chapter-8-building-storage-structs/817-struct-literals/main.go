package main

import (
	"fmt"
	"github.com/lethang7794/magazine"
)

func main() {
	s1 := magazine.Subscriber{ // Literal for a Subscriber struct.
		Name:   "Aman Singh",
		Rate:   4.99,
		Active: true,
	}

	fmt.Println(s1)

	s2 := magazine.Subscriber{ // Omit some fields from the curly braces.
		Name: "Harry",
	}
	fmt.Println(s2)

	s3 := magazine.Subscriber{} // Omit all fields
	fmt.Println(s3)             // Omitted fields will be set to zero value.
}
