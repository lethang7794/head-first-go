package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	// for index, value := range notes { // index declared but not used
	// 	fmt.Println(value)
	// }

	for _, value := range notes {
		fmt.Println(value)
	}

	for index, _ := range notes {
		fmt.Println(index)
	}

	for index := range notes { // simplify range expression
		fmt.Println(index)
	}
}
