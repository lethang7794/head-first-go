package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

// Needs to be a POINTER RECEIVER, so original value can be updated
func (d *Date) SetYear(year int) {
	d.Year = year // Updates original value, not a copy
	// ~ (*d).Year = year
}

func main() {
	var date Date
	date.SetYear(2022)
	fmt.Println(date)
}
