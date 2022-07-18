package main

import (
	"fmt"
	"time"
)

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	// Assigns values to struct fields
	myStruct.number = 3
	myStruct.word = "pie"
	myStruct.toggle = true

	// Retrieve values from struct fields
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)

	// "Dot operator" indicates fields "belong to" a struct
	// This works for both assigning values and retrieving them.

	fmt.Println("hi")
	// Println function "belong to" fmt package

	var myTime time.Time
	myTime.Year()
	// Year method "belong to" a "Time" value
}
