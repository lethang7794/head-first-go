package main

import "fmt"

func calmDown() {
	fmt.Println(recover())
}

func main() {
	defer calmDown()
	panic("oh no")
}
