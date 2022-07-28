package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk")
}

type Robot string // Declare a new Robot type

func (r Robot) MakeSound() { // Robot satisfies the NoiseMaker interface
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() { // Robot has an addition method
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound() // OK. Part of NoiseMaker interface.
	//n.Walk() // Not OK. Not part of NoiseMaker!
	//n.Walk undefined (type NoiseMaker has no field or method Walk)
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	play(toy)

	toy = Robot("Botco Ambler")
	play(toy)
}
