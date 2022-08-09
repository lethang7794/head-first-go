package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(fn func()) {
	fn()
	fn()
}

func main() {
	twice(sayHi)
	twice(sayBye)
}
