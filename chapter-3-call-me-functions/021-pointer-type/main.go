package main

import (
	"fmt"
	"reflect"
)

func main() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)

	var (
		myInt   int
		myFloat float64
		myBool  bool
	)
	fmt.Println(&myInt, reflect.TypeOf(&myInt))
	fmt.Println(&myFloat, reflect.TypeOf(&myFloat))
	fmt.Println(&myBool, reflect.TypeOf(&myBool))

	var (
		myIntPointer   *int
		myFloatPointer *float64
		myBoolPointer  *bool
	)

	fmt.Println(myIntPointer)
	fmt.Println(myFloatPointer)
	fmt.Println(myBoolPointer)

	myIntPointer = &myInt
	myFloatPointer = &myFloat
	myBoolPointer = &myBool

	fmt.Println(myIntPointer)
	fmt.Println(myFloatPointer)
	fmt.Println(myBoolPointer)
}
