package main

import (
	"fmt"
	"github.com/lethang7794/magazine"
)

func main() {
	var address magazine.Address
	address.Street = "123 Oak St"
	address.City = "Ohama"
	address.State = "NE"
	address.PostalCode = "68111"

	fmt.Println(address)
}
