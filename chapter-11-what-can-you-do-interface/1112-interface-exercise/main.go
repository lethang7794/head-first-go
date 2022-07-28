package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}

func (c Car) Brake() {
	fmt.Println("Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}

func (t Truck) Brake() {
	fmt.Println("Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func main() {
	var vehicle Vehicle = Car("Toyota Yarvic")
	vehicle.Accelerate()
	vehicle.Steer("left")

	vehicle = Truck("Fnord F180")
	vehicle.Brake()
	vehicle.Steer("right")
}
