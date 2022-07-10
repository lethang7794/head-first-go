package main

import "fmt"

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, v := range numbers {
		sum += v
	}
	return sum / float64(len(numbers))
}

func main() {
	fmt.Println(average(1, 2, 3, 4, 5))
	fmt.Println(average(11.11, 22.22, 33.33))
	fmt.Println(average(100, 50))
	fmt.Println(average(90.7, 89.7, 98.5, 92.3))
}
