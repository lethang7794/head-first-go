package main

import "fmt"

func calmDown2() {
	p := recover() // Returns an interface{} value
	//fmt.Println(p.Error()) // Even though the underlying "error" value has an Error method, the interface{} value doesn't!
	fmt.Println(p)
}

func main() {
	calmDown2()
	err := fmt.Errorf("there's an error")
	panic(err) // Instead of a string, pass an error value to "panic"
}
