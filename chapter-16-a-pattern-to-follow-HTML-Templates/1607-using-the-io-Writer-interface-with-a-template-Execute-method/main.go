package main

import "fmt"

func main() {
	fmt.Println("What is os.Stdout value?")
	// os.Stdout value is part of the os package.
	// Stdout stands for "standard output".
	// It's acts like a file, but any data written to it is
	// output to the terminal instead of being saved to disk.

	fmt.Println("How can an http.ResponseWriter and os.Stdout both be valid values to pass to a Template's Execute method?")
	// func (t *Template) Execute(wr io.Writer, data any) error

	// type Writer interface {
	//	Write(p []byte) (n int, err error)
	//}

	// Both has the Write(p []byte) (n int, err error) method
}
