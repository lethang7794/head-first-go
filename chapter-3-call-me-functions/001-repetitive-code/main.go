package main

import "fmt"

func main() {
	var width, height, area float64

	width = 4.2
	height = 3.0
	area = width * height
	fmt.Println(area/10.0, "liters needed")

	width = 5.2
	height = 3.5
	area = width * height
	fmt.Println(area/10.0, "liters needed")
}
