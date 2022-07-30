package main

import "fmt"

func recurses() {
	fmt.Println("Oh, no, I'm stuck!")
	recurses()
}

func main() {
	recurses()
}
