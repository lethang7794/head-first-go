package main

import "fmt"

var fuel float64 = 10
var _ = fuel // Does that represent 10 gallons or 10 liters

// Define two new types, each with an underlying type of float64
type Liters float64
type Gallons float64

func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10) // Convert a float64 to Gallons
	busFuel = Liters(10)  // Convert a float64 to Liters
	fmt.Println(carFuel, busFuel)

	// carFuel = busFuel // Cannot use 'busFuel' (type Liters) as the type Gallons

	carFuel = Gallons(busFuel) // 10 liters does NOT equal 10 gallons

	carFuel = Gallons(busFuel * 0.264) // Convert from Liters to Gallons
	busFuel = Liters(carFuel * 3.785)  // Convert from Gallons to Liters
	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)
}
