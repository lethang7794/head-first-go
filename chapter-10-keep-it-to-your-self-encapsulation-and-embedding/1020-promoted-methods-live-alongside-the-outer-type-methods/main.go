package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/calendar"
	"log"
)

func main() {
	event := calendar.Event2{}

	err := event.SetTitle("Girlfriend's birthday") // Defined on Event2 itself
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(2022) // Promoted from Date
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(7) // Promoted from Date
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(25) // Promoted from Date
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Everything is good! 😄")
	fmt.Println(event.Title()) // Defined on Event2 itself
	fmt.Println(event.Year())  // Promoted from Date
	fmt.Println(event.Month()) // Promoted from Date
	fmt.Println(event.Day())   // Promoted from Date
}
