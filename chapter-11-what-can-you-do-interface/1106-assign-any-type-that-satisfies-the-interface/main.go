package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() { // Whistle has a MakeSound method.
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() { // Horn also has a MakeSound method.
	fmt.Println("Honk")
}

type NoiseMaker interface { // Represents any type with a MakeSound method.
	MakeSound()
}

func play(n NoiseMaker) { // Declare function parameters with interface types as well
	n.MakeSound()
}

func main() {
	var toy NoiseMaker // Declare a NoiseMaker variable.

	toy = Whistle("Toyco Canary") // Assign a value of a type that satisfies NoiseMaker to the variable.
	toy.MakeSound()

	toy = Horn("Toyco Blaster") // Assign a value of another type that satisfies NoiseMaker to the variable.
	toy.MakeSound()

	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))
}
