package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, v := range numbers {
		sum += v
	}
	return sum / float64(len(numbers))
}

func main() {
	arguments := os.Args[1:]
	var numbers []float64
	for _, arg := range arguments {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	// fmt.Println(average(numbers))    // cannot use numbers (variable of type []float64) as float64 value in argument to average
	fmt.Println(average(numbers...)) // cannot use numbers (variable
}
