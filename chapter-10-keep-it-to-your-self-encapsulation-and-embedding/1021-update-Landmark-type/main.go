package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/geo2"
	"log"
)

func main() {
	location := geo2.Landmark{}
	err := location.SetName("The Googleplex")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
