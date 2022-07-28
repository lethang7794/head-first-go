package main

import "fmt"

type Gallons float64
type Liters float64
type Milliliters float64

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(Milliliters(12.09248342))
	// All three values look identical!

	fmt.Printf("%0.2f gal\n", Gallons(12.09248342))
	fmt.Printf("%0.2f L\n", Liters(12.09248342))
	fmt.Printf("%0.2f mL\n", Milliliters(12.09248342))
	// Format the numbers and add abbreviations.

	coffeePot := CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())
	fmt.Println(coffeePot)
	fmt.Printf("%s", coffeePot)
	// The return value of String is used in the output.
}
