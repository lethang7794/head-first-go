package main

import (
	"fmt"
	"log"

	"github.com/lethang7794/head-first-go/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := len(numbers)
	fmt.Printf("Average: %0.2f\n", sum/float64(sampleCount))
}
