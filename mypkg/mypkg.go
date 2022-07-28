package mypkg

import "fmt"

// Declare an interface type.
type MyInterface interface {
	MethodWithoutParameters()      // A type satisfies this interface if it has this method...
	MethodWithParameter(float64)   // and this method (with a float64 parameter)...
	MethodWithReturnValue() string // and this method (with a string return value).
}

// Declare a type. We'll make it satisfy MyInterface
type MyType int

func (m MyType) MethodWithoutParameters() { // First required method.
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWithParameter(f float64) { // Second required method (with a float64 parameter).
	fmt.Println("MethodWithParameter called with", f)
}

func (m MyType) MethodWithReturnValue() string { // Third required method (with a string return value).
	fmt.Println("MethodWithReturnValue called")
	return "Hi from MethodWithReturnValue"
}

// A type can still satisfy an interface even if it has methods that aren't part of the interface.
func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
