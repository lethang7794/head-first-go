package main

// An INTERFACE is a set of methods
// that certain values are expected to have.

// Define an interface type using interface keyword.
type myInterface interface {
	methodWithoutParameters()
	methodWithParameter(float64)
	methodWithReturnValue() string
}

type _ myInterface

// Any type that has all the methods listed in an interface definition
// is said to SATISFY that interface.

// A type that satisfies an interface can be used anywhere that interface is called for.

func main() {

}
