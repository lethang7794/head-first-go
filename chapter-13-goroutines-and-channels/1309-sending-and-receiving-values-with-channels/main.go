package main

import "fmt"

func greeting(myChannel chan string) {
	fmt.Println("Sending greeting")
	myChannel <- "hi"
}

func main() {
	// myChannel := make(chan float64)
	// myChannel <- 3.14 // Send 3.14 value to the myChannel
	// fmt.Println(<-myChannel) // Receive value from myChannel - THIS WON'T WORK.

	myStringChannel := make(chan string)
	go greeting(myStringChannel)

	fmt.Println("Receiving greeting")
	fmt.Println(<-myStringChannel)     // Directly, or ...
	receivedValue := <-myStringChannel // ...store the received value in a variable instead.
	fmt.Println(receivedValue)
}
