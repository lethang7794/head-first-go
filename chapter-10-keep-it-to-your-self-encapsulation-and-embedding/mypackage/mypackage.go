package mypackage

import "fmt"

type EmbeddedType string

type MyType struct {
	EmbeddedType
}

func (e EmbeddedType) ExportedMethod() {
	fmt.Println("Hi from ExportedMethod on EmbeddedType")
}

func (e EmbeddedType) unexportedMethod() {
	fmt.Println("Hi from unexportedMethod on EmbeddedType")
}
