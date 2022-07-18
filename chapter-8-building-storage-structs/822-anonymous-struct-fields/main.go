package main

import (
	"fmt"
	"github.com/lethang7794/magazine"
)

func main() {
	author := magazine.Author{
		Book: "How to be fun?",
		Address: magazine.Address{
			Street:     "789 Big St",
			City:       "Omaha",
			State:      "OR",
			PostalCode: "9722",
		},
	}

	fmt.Println("author.Address.City", author.Address.City)
}
