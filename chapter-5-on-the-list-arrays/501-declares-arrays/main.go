package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 42
	numbers[2] = 108
	var letters = [3]string{"a", "b", "c"}

	fmt.Println(numbers[0]) // 42
	fmt.Println(numbers[1]) // 0
	fmt.Println(numbers[2]) // 108
	fmt.Println(letters[2]) // c
	fmt.Println(letters[0]) // a
	fmt.Println(letters[1]) // b
}
