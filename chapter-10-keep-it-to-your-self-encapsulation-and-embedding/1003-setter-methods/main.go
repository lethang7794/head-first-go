package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d Date) SetYear(year int) { // Accepts the value the field should be set to.
	d.Year = year // Set the struct field.
	fmt.Println(d.Year)
}

func main() {
	var date Date      // Create a Date
	date.SetYear(2022) // Set its Year field via the method
	fmt.Println(date)  // Print the year field
	// ! Year is still set to its zero value!
}
