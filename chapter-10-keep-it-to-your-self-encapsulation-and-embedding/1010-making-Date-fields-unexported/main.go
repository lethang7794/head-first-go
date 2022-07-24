package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/calendar"
)

func main() {
	date := calendar.Date2{}

	// We can't directly set field values anymore
	//date.year = 2022 // date.year undefined (type calendar.Date2 has no field or method year)
	//date.month = 14
	//date.day = 50

	// Even through struct literal
	date = calendar.Date2{
		//year:  0, // unknown field 'year' in struct literal of type calendar.Date2
		//month: 0,
		//day:   -2,
	}
	fmt.Println(date)
}
