package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	
	index := 1
	fmt.Println(index, notes[index])

	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}
	
	// fmt.Println(notes[7]) // invalid argument: array index 7 out of bounds [0:7])

	for i := 0; i <= 7; i++ {
		// fmt.Println(i, notes[i]) // panic: runtime error: index out of range [7] with length 7
	}
}
