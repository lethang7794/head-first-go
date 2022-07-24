package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) {
	d.Year = year
}

func (d *Date) SetMonth(month int) {
	d.Month = month
}

func (d *Date) SetDay(day int) {
	d.Day = day
}

func main() {
	var date Date
	date.SetYear(2022)
	date.SetMonth(7)
	date.SetDay(22)
	fmt.Println(date)

	var date2 Date
	date2.SetYear(0)   // invalid year
	date2.SetMonth(15) // invalid month
	date2.SetDay(50)   // invalid day
	fmt.Println(date2)
}
