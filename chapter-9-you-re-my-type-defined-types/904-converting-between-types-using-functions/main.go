package main

import "fmt"

type Liters float64
type Gallons float64

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	// carFuel += Liters(8.0) // Invalid operation: carFuel += Liters(8.0) (mismatched types Gallons and Liters)
	// busFuel += Gallons(30.0) // Invalid operation: busFuel += Gallons(30.0) (mismatched types Liters and Gallons)

	carFuel += ToGallons(Liters(40))
	busFuel += ToLiters(Gallons(30))

	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
}
