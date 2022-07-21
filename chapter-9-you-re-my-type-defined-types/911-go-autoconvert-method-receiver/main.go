package main

import "fmt"

type MyType string

func (m MyType) valueMethod() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := MyType("MyType value")
	pointer := &value

	value.valueMethod()
	pointer.pointerMethod()

	// Value automatically converted to pointer
	value.pointerMethod()

	// Pointer automatically converted to value
	pointer.valueMethod()

	MyType("MyType value").valueMethod()

	// To call a method that requires a pointer receiver,
	// we have to be able to get a pointer to the value
	//MyType("MyType value").pointerMethod() // Cannot call a pointer method on 'MyType("MyType value")'
}
