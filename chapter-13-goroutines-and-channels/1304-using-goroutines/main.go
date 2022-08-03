package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func c() {
	for i := 0; i < 50; i++ {
		fmt.Print("c")
	}
}

func d() {
	for i := 0; i < 50; i++ {
		fmt.Print("d")
	}
}

func main() {
	a()
	b()
	go c() // Where is the output of c()...
	go d() // ...and d()?

	fmt.Println("end main()")

	// Go program stops running as soon as the main goroutine ends,
	// EVEN if other goroutines are still running.
	// -> main function completes before the code in c and d function have a chance to run.

	// We need to keep the main goroutine running until the goroutine for the c and d function can finish.
	// For now: we'll just pause the main goroutine for a set amount of time so the other goroutines can run.
	time.Sleep(time.Second)
}
