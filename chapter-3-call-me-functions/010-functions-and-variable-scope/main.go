package main

import "fmt"

var meterPerLiter float64

func paintNeeded(width, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/meterPerLiter)
}

func main() {
	meterPerLiter = 10.0
	paintNeeded(4.2, 3.0)
	// fmt.Println(area) // ! undeclared name: area
}
