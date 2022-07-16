package main

import (
	"fmt"
)

func main() {
	_ = []string{"Amber Graham", "Brian Martin", "Carlos Diaz"}
	_ = []int{111, 222, 333}
	// We need to search though every name to know how many votes a name has

	// Declare a map variable
	var myMap map[string]float64
	fmt.Printf("myMap %v %#v\n", myMap, myMap)

	// Declare a map variablex
	var ranks map[string]int
	// Actually create a map
	ranks = make(map[string]int)

	// Create a map and declare a variable to hold it
	scores := make(map[string]float64)
	fmt.Println(scores)

	// Assigns values to a map
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3

	// Get values from a map
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["bronze"])

	// A map with both keys and values of type string
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])

	// A map with integers as keys and booleans as values
	isPrime := make(map[int]bool)
	isPrime[1] = false
	isPrime[2] = true
	isPrime[3] = true
	isPrime[4] = false
	isPrime[5] = true
	fmt.Println(isPrime)
}
