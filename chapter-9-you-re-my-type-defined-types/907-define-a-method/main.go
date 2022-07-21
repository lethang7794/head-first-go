package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Printf("Hi from %v, %v\n", m, &m)
}

func (m *MyType) sayHello() {
	fmt.Printf("Hello from %v, %v\n", m, *m)
}

func main() {
	value := MyType("a MyType value")
	value.sayHi()
	value.sayHi()

	value.sayHello()
	value.sayHello()
}
