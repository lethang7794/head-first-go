package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/mypkg"
)

func main() {
	// Declare a variable using MyInterface type
	var value mypkg.MyInterface

	// Values of MyType satisfy myInterface,
	// so we can assign this value to a variable with a type of MyInterface
	value = mypkg.MyType(5)

	// We can call any method that's part of MyInterface
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
