package main

import (
	"time"
)

func main() {
	// size := go responseSize("https://example.com") // syntax error: unexpected go, expecting expression

	// size = go responseSize("https://golang.org")
	// The compiler stops we from attempting to
	// get a return value from a function called with a go statement.

	// size = go responseSize("https://golang.org/doc")
	// We're saying: "go run this, I'm not going to wait."
	// But we also immediately attempt to use the return value (which may not ready yet).

	// There is a way to communicate between goroutines: CHANNELS.
	// - Not only do channels allow we to send values from one goroutine to another,
	// - But also they ensure the sending goroutine has SENT the value BEFORE
	// the receiving goroutine attempts to use it.

	// Each channel only carries values of a particular type

	// Declare a variable to hold a channel carries...
	var myIntChannel chan int         // ...int values
	var myFloat64Channel chan float64 // ...float64 values

	// Create the channels
	myIntChannel = make(chan int)
	myFloat64Channel = make(chan float64)

	// Create a channel and declare a variable at once
	myBoolChannel := make(chan bool)

	_, _, _ = myIntChannel, myFloat64Channel, myBoolChannel

	time.Sleep(5 * time.Second) // Sleep for 5 seconds.
}
