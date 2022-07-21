package main

import "fmt"

func main() {
	// A defined type supports all the same operations as its underlying type.

	// Types based on float64, for example,
	// support arithmetic operators like +, -, *, /,
	// as well as comparison operators like ==, > and <
	type Liters float64
	type Gallons float64

	fmt.Println(Liters(1.2) + Liters(3.4))
	fmt.Println(Gallons(5.5) - Gallons(2.2))
	fmt.Println(Liters(2) * Liters(3.5))
	fmt.Println(Gallons(2.2) / Gallons(1.1))
	fmt.Println(Liters(1.2) < Liters(3.4))
	fmt.Println(Liters(1.2) > Liters(3.4))

	// A type based on an underlying type of string,
	// however, would support +, ==, > and <,
	// but not -, because - is not a valid operator for strings
	type Title string
	fmt.Println(Title("Alien") == Title("Alien"))
	fmt.Println(Title("Alien") < Title("Alien"))
	fmt.Println(Title("Alien") > Title("Alien"))
	fmt.Println(Title("Alien") + "s")
	//fmt.Println(Title("Jaws 2") - " 2") // Invalid operation: Title("Jaws 2") - " 2" (the operator - is not defined on string)

	// A defined type can be used in operator together with literal values.
	fmt.Println(Liters(1.2) + 3.4)
	fmt.Println(Gallons(5.5) - 2.2)
	fmt.Println(Liters(1.2) == 1.2)
	fmt.Println(Gallons(1.2) < 3.4)

	// But defined types cannot be used in operations together with values of a different type,
	// even if the other type has the same underlying type.
	// fmt.Println(Liters(1.2) + Gallons(3.4)) // Invalid operation: Liters(1.2) + Gallons(3.4) (mismatched types Liters and Gallons)
	// fmt.Println(Gallons(1.2) == Liters(1.2)) // Invalid operation: Gallons(1.2) == Liters(1.2) (mismatched types Gallons and Liters)
}
