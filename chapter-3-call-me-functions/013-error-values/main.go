package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("height can't be negative")
	fmt.Println(err.Error())
	// log.Fatal(err)

	err2 := fmt.Errorf("a height of %0.2f is invalid", -2.33333)
	fmt.Println(err2.Error())
}
