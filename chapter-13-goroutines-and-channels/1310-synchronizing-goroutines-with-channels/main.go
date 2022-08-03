package main

import "fmt"

// Channels ensure the sending goroutine HAS SENT the value,
// BEFORE the receiving channel attempts to USE it.

// Channels do this by BLOCKING:
// pausing all further operations in the current goroutine
// until another goroutine executes a receive operation on the same channel.

// And VICE VERSA,
// a receive operation blocks the receiving goroutine
// until another goroutine executes a send operation on the same channel.

// => This behavior allows goroutines to SYNCHRONIZE their actions,
// - tha is, to coordinate their timing.

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	// main goroutine become the orchestrator of the abc and def goroutines,
	// allowing them to proceed only when it (main) is ready to read the values they're sending.

	// Create two channels.
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Pass each channel to a function running in a new goroutine.
	go abc(channel1)
	go def(channel2)

	// Receive and print values from the channels, IN ORDER.
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
}
