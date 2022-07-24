package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/geo2"
	"log"
)

func main() {
	coordinates := geo2.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-1122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}
