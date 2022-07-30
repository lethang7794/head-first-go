package main

import (
	"fmt"
	"log"
)

func Socialize() error {
	defer fmt.Println("Goodbye") // Defer printing "Goodbye"
	fmt.Print("Hello!")
	return fmt.Errorf("i don't want to talk") // Return an error

	fmt.Println("Nice weather, eh?") // This code won't run
	return nil
}

func main() {
	err := Socialize()
	if err != nil {
		log.Fatal()
	}
}
