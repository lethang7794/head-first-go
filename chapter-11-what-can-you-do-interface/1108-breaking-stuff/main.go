package main

import "fmt"

type Appliance interface {
	TurnOn()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}

type CoffeePot string

func (c CoffeePot) TurnOn() error {
	fmt.Println("Powering Up")
	return nil
}
func (c CoffeePot) Brew() {
	fmt.Println("Heating Up")
}

func main() {
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()

	device = CoffeePot("LuxBrew")
	device.TurnOn()

	//device.Brew() // device.Brew undefined (type Appliance has no field or method Brew)
}
