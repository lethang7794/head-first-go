package main

import (
	"fmt"
	"github.com/lethang7794/magazine"
)

func main() {
	author1 := magazine.Author{
		Book: "Golang for newbie",
	}

	// The anonymous field is EMBEDDED inside the outer struct.
	// Fields of the embedded struct are PROMOTED to the outer struct.

	// (Address type) is EMBEDDED inside (Author type).
	// Address's fields can be assigned directly as if there is no Address type at all.
	author1.Street = "123 Oak St"
	author1.City = "Omaha"
	author1.State = "NE"
	author1.PostalCode = "68111"

	// And retrieve as well.
	fmt.Println("Street:", author1.Street)
	fmt.Println("City:", author1.City)
	fmt.Println("State:", author1.State)
	fmt.Println("PostalCode:", author1.PostalCode)

	fmt.Printf("author1: %#v", author1)
}
