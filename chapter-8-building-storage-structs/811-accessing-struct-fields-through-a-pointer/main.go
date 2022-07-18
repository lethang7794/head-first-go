package main

import "fmt"

func main() {
	var value int = 2         // Create a value
	var pointer *int = &value // Get a pointer to the value

	fmt.Println("pointer:", pointer)   // Print the pointer, not the value
	fmt.Println("*pointer:", *pointer) // Print the value at the pointer

	type myStruct struct {
		myField      int
		myOtherField string
	}
	var myValue myStruct
	myValue.myField = 3
	var myPointer *myStruct = &myValue
	fmt.Println("myPointer:", myPointer)

	// Attempt to get the struct value at the pointer
	//fmt.Println("myPointer.myField:", *myPointer.myField) // ! Invalid indirect of 'myPointer.myField' (type 'int')

	// Get the struct value at the pointer, THEN access the struct field
	fmt.Println("(*myPointer).myField:", (*myPointer).myField)

	// ACCESS the struct field THROUGH the pointer
	fmt.Println("myPointer.myField:", myPointer.myField)

	// ASSIGN the struct field  THROUGH a pointer works too
	myPointer.myOtherField = "This field value is assigned through a pointer as well"
	fmt.Println("myPointer.myOtherField", myPointer.myOtherField)
}
