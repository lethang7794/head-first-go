package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	// The Event struct Title field is exported, we can still access it directly
	event.Title = "A very very very very very long title"
	fmt.Println(event.Title)

	event2 := calendar.Event2{}
	//event2.Title = "A title" // This is prevented
	err := event2.SetTitle("A very very very very very long title")
	if err != nil {
		log.Fatal(err)
	}
}
