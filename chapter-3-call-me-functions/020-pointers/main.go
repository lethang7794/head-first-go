package main

import "fmt"

func main() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)

	var (
		myInt   int
		myFloat float64
		myBool  bool
	)

	fmt.Println(&myInt)
	fmt.Println(&myFloat)
	fmt.Println(&myBool)
}
