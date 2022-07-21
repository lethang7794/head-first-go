package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (m Milliliters) ToLiters() Liters {x
	return Liters(m / 1000)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Liters {
	return Liters(g * 3785.41)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals to %0.3f gallons\n", soda, soda.ToGallons())
	fmt.Printf("%0.3f liters equals to %0.3f milliliters\n", soda, soda.ToMilliliters())

	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals to %0.3f gallons\n", water, water.ToGallons())
	fmt.Printf("%0.3f milliliters equals to %0.3f liters\n", water, water.ToLiters())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals to %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals to %0.3f milliliters\n", milk, milk.ToMilliliters())
}
