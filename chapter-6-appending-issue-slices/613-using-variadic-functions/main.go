package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64

	for _, v := range numbers {
		if min <= v && v <= max {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	fmt.Println(maximum(1, 2, 3, 4, 5))
	fmt.Println(maximum(9, 8, 7, 6, 5))
	fmt.Println(maximum(-5, -4, -3, -2, -1))
	fmt.Println(maximum(-1, -2, -3, -4, -5))
	fmt.Println(maximum(-2, -1, 0, 1, 2))
	fmt.Println(maximum(-2.2, -1.1, 0, 1.1, 2.2))
	fmt.Println(maximum())

	fmt.Println(inRange(10, 20, 1, 2, 3, 10, 11, 19, 20, 21, 22))
	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))
}
