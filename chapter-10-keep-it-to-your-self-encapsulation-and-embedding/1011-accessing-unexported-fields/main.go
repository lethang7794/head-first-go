package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/calendar"
	"log"
)

func main() {
	date := calendar.Date2{}

	// But we still can update the unexported field value
	// ... through the setters
	err := date.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(22)
	if err != nil {
		log.Fatal(err)
	}

	// We can see the unexported fields
	fmt.Println(date)
	fmt.Printf("%#v\n", date)

	// But how do we get those values back out?
}
