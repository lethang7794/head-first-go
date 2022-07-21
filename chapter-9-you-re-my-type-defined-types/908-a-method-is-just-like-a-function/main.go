package main

import "fmt"

type MyType string

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) ExportedMethod() {

}

func (m MyType) unExportedMethod() {

}

func main() {
	value := MyType("MyType Value")
	value.MethodWithParameters(4, true)
}
