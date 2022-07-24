package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/calendar"
)

func main() {
	event := calendar.Event{}

	// Unexported Date fields aren't promoted to the Event
	//event.month = 5 // event.month undefined (type calendar.Event has no field or method month, but does have Month)

	// Can't access Date2 fields directly on the Date2
	//event.Date2.year // event.Date2.year undefined (type calendar.Date2 has no field or method year, but does have Year)

	fmt.Println(event)
}
