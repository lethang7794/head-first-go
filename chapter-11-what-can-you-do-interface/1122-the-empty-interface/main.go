package main

import "fmt"

type Anything interface{} // An EMPTY INTERFACE, it doesn't require any methods at all.
// ANY type would satisfy this interface

type SomeThing string

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("zZz")
}

func AcceptAnything(thing interface{}) { // Accepts a parameter with the empty interface as its type
	fmt.Println(thing)
	//thing.SomeFancyMethod() // Can't do this
	// And empty interface doesn't have any methods.

	whistle, ok := thing.(Whistle) // We can do a type assertion for value of empty interface
	if ok {
		whistle.MakeSound()
	}
}

// But it's better to write a new function
func AcceptWhistle(whistle Whistle) {
	whistle.MakeSound() // No type conversion needed
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything("A string")
	AcceptAnything(true)
	AcceptAnything(SomeThing("some thing"))

	AcceptAnything(Whistle(""))

	AcceptWhistle(Whistle(""))
}
