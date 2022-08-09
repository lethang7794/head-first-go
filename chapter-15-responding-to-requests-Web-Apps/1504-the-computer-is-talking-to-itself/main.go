package main

import "fmt"

func main() {
	fmt.Println("What are the differences between localhost:8080 and 0.0.0.0:8080")
	// localhost:8080
	// ---------:----
	// Host     :Port

	// localhost: Establish a connection FROM your computer TO ITSELF.

	fmt.Println("What is a port?")
	// A port is a numbered network communication channel
	// that an application can listen for messages on.
	// E.g:
	// - 80 --> http
	// - 443 -> https
}
