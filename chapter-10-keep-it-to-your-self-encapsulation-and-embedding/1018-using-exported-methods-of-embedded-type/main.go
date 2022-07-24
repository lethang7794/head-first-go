package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/calendar"
	"log"
)

func main() {
	event := calendar.Event{}

	//event.month = 7  // Can't do this
	//event.Date2.year = 2022 // Neither this

	err := event.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(23)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
