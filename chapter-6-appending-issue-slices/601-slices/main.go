package main

import "fmt"

func printSliceString(slice []string) {
	fmt.Printf("%#v %v %T\n", slice, slice, slice)
}

func main() {
	var mySlice []string
	printSliceString(mySlice)

	// mySlice[0] = "Will this be the first element?" // panic: runtime error: index out of range [0] with length 0

	mySlice = append(mySlice, "first element")
	printSliceString(mySlice)

	mySlice[0] = "Will this override the first element?"
	printSliceString(mySlice)

	// mySlice[1] = "Will this override the second element?" // panic: runtime error: index out of range [1] with length 1
}
