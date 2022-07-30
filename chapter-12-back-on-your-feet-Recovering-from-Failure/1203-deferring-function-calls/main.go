package main

import "fmt"

func Socialize() {
	defer fmt.Println("Goodbye!") // This function is deferred until after Socialize exits!
	fmt.Println("Hello!")
	fmt.Println("Nice weather, eh?")
}
func main() {
	Socialize()
}
