package main

import "fmt"

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func main() {
	myChannel := make(chan string)

	// greeting(myChannel) // fatal error: all goroutines are asleep - deadlock!

	go greeting(myChannel)
	fmt.Println(<-myChannel)

	// myChannel <- "hello" // fatal error: all goroutines are asleep - deadlock!
}
