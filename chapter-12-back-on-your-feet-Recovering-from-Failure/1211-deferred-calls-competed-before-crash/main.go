package main

import "fmt"

func main() {
	one()
}

func one() {
	defer fmt.Println("deferred in one()")
	two()
}

func two() {
	defer fmt.Println("deferred in two()")
	panic("Let's see what's been deferred!")
}
