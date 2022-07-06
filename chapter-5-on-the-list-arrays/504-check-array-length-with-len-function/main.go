package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(len(notes))

	primes := [5]int{}
	fmt.Println(len(primes))

	emptyArr := []int{}
	fmt.Println(len(emptyArr))

	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}

	for i := 0; i <= len(notes); i++ {
		// fmt.Println(i, notes[i]) // panic: runtime error: index out of range [7] with length 7
	}
}
