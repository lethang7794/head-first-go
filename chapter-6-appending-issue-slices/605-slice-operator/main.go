package main

import "fmt"

func main() {
	// When we create a slice, there is an UNDERLYING ARRAY created automatically.
	// The slice is just a view into the UNDERLYING ARRAY.
	// Although we can't access the UNDERLYING ARRAY directly (without thought the slice).
	aSlice := make([]int, 10)

	_ = aSlice // We know that aSlice is not used

	// We can create an array ourselves,
	// then create a slice based on it thought the SLICE OPERATOR
	underlyingArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	firstTwo := underlyingArray[0:2]  // START from 0; STOP BEFORE 2
	nextThree := underlyingArray[2:5] // START from 2; STOP BEFORE 5

	fmt.Println(firstTwo)
	fmt.Println(nextThree)

	oneToFour := underlyingArray[1:5]    // START from 1; STOP BEFORE 5 ==> 5 - 1 = 4 elements
	threeToNine := underlyingArray[3:10] // START from 3; STOP BEFORE 10 ==> 10 - 3 = 7 elements

	fmt.Println(oneToFour)
	fmt.Println(threeToNine)

	first := underlyingArray[:1]
	fmt.Println(first)

	allExceptFirst := underlyingArray[1:]
	fmt.Println(allExceptFirst)

	last := underlyingArray[len(underlyingArray)-1:]
	fmt.Println(last)
}
