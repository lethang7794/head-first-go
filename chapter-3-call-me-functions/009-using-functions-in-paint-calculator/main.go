package main

import "fmt"

func paintNeeded(width, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)
}
