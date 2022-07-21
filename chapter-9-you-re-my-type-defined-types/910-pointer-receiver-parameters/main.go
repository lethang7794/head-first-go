package main

import "fmt"

type Number int

// Receiver parameter are treated no differently than ordinary parameters.
// A receiver paramter receives a COPY of the receiver value.
func (n Number) Double() {
	n *= 2 // Changing the COPY, not the original.
}

// To modify the receiver, we need to use a pointer receiver.
func (n *Number) DoublePointerReceiver() {
	*n *= 2
}

func main() {
	number := Number(4)
	fmt.Println("Original value of number:", number)
	number.Double()
	fmt.Println("Number after call Double:", number)

	// If you call a method requiring a value receiver,
	// Go automatically get the value at the pointer and pass to the method
	number.DoublePointerReceiver()
	fmt.Println("Number after call DoublePointerReceiver:", number)
}
