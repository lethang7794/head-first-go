package main

import "fmt"

type Part struct {
	description string
	count       int
}

func showInfo(p Part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func minimumOrder(description string) Part {
	var p Part
	p.description = description
	p.count = 100
	return p
}

func main() {
	var bolts Part
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts)

	p := minimumOrder("Hex bolts")
	fmt.Println(p)
}
