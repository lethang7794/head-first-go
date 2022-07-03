package main

import "fmt"

func main() {
	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
	fmt.Printf("That will be $%f please.\n", 0.23*5)

	fmt.Printf("%v %v %v", "", "\t", "\n")
	fmt.Printf("%#v %#v %#v", "", "\t", "\n")
}
