package main

import "fmt"

func main() {
	myFloat := 3.14
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	myBool := true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)

	myInt := 4
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)

	*myIntPointer = 8
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}
