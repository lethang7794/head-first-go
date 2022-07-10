package main

import "fmt"

func main() {
	// The slice's values are known in advance
	primes := []int{2, 3, 5, 7, 11}
	fmt.Println(primes)

	// Multi-line slice literal
	notes := []string{
		"do",
		"re",
		"mi",
		"fa",
		"so",
		"la",
		"si",
	}
	fmt.Println(notes)
}
