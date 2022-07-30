package main

import "fmt"

func count(start int, end int) {
	fmt.Printf("count(%d, %d) called\n", start, end)
	fmt.Println(start) // Print the current starting number.
	if start < end {   // If we haven't reached the ending number...
		count(start+1, end) // ...the "count" function calls itself, with a starting number 1 higher than before
	}
	fmt.Printf("Returning from count(%d, %d) call\n", start, end)
}

func main() {
	count(1, 10000)
}
