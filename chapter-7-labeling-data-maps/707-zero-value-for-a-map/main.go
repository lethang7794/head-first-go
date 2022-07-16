package main

import "fmt"

func main() {
	var nilMap map[int]string
	fmt.Printf("nilMap: %#v\n", nilMap)
	// nilMap[3] = "three" // panic: assignment to entry in nil map

	var myMap map[int]string = make(map[int]string)
	myMap[4] = "four"
	fmt.Printf("myMap: %#v\n", myMap)
}
