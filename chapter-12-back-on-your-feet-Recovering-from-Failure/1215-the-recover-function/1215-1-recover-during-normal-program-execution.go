package main

import "fmt"

func main() {
	fmt.Println(recover()) // In a program that isn't panicking, recover() does nothing, and return nil.
}
