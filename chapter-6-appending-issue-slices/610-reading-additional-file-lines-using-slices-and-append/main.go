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

	var sum float64
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Average %0.2f", sum/float64(len(numbers)))
}
