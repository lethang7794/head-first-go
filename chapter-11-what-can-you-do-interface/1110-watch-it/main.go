package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func main() {
	s := Switch("off")
	// var t Toggleable = s // cannot use s (variable of type Switch) as type Toggleable in variable declaration:
	// 							Switch does not implement Toggleable (toggle method has pointer receiver)

	var t Toggleable = &s // Assign a pointer to a Switch instead
	t.toggle()
	t.toggle()
}
