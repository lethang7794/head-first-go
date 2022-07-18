package main

import "fmt"

func main() {
	// Define the struct type for the "subscriber1" variable
	var subscriber1 struct {
		name   string
		rate   float64
		active bool
	}

	subscriber1.name = "Aman Singh"
	subscriber1.rate = 4.99
	subscriber1.active = true

	fmt.Println(subscriber1)

	// Define an identical type all over again for the "subscriber2" variable
	var subscriber2 struct {
		name   string
		rate   float64
		active bool
	}

	subscriber2.name = "Harry Potter"
	subscriber2.rate = 1.99
	subscriber2.active = false

	fmt.Println(subscriber2)

	// Define a type name "car"
	type car struct {
		name     string
		topSpeed float64
	}

	// Declare a variable of type "car"
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323

	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	// Define a type name "part"
	type part struct {
		description string
		count       int
	}

	// Declare a variable of type "part"
	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24

	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)
}
