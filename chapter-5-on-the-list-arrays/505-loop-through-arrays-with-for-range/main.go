package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	for index, value := range notes {
		fmt.Println(index, value)
	}

}
