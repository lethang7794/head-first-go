package main

import (
	"fmt"
	"github.com/lethang7794/magazine"
)

func main() {
	// Create an Address value and populate its fields.
	address := magazine.Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "68111",
	}

	// Create the Subscriber struct that the Address will belong to.
	subscriber := magazine.Subscriber{
		Name: "Aman Singh",
	}
	// Set the HomeAddress field
	subscriber.HomeAddress = address

	// Print the HomeAddress field
	fmt.Println(subscriber.HomeAddress)

	subscriber2 := magazine.Subscriber{}
	fmt.Printf("%#v\n", subscriber2.HomeAddress)
	// The HomeAddress field is already set as a new Address struct
	// Each of the Address struct's fields is set to its zero value

	// "Chain" dot operators together so we can access the fields of Address struct
	fmt.Printf("subscriber.HomeAddress.City: %#v\n", subscriber.HomeAddress.City)

	// This works for both the assigning values to the inner struct's fields ...
	subscriber2.HomeAddress.PostalCode = "68111"

	// ...and for retrieving those values again later
	fmt.Printf("PostalCode: %#v\n", subscriber2.HomeAddress.PostalCode)

	subscriber3 := magazine.Subscriber{Name: "Harry Potter"}
	subscriber3.HomeAddress.Street = "123 Oak St"
	subscriber3.HomeAddress.City = "Omaha"
	subscriber3.HomeAddress.State = "NE"
	subscriber3.HomeAddress.PostalCode = "68111"

	fmt.Printf("%#v\n", subscriber3)

	employee := magazine.Employee{
		Name: "Hermione Granger",
		HomeAddress: magazine.Address{
			Street:     "456 Elm St",
			City:       "Portland",
			State:      "OR",
			PostalCode: "97222",
		},
	}

	fmt.Printf("%#v\n", employee)

}
