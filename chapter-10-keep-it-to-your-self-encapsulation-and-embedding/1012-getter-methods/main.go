package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/calendar"
	"log"
)

func main() {
	date := calendar.Date2{}

	err := date.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(23)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
