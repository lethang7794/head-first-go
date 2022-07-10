package main

import "fmt"

func twoInts(first int, second int) {
	fmt.Println(first, second)
}

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func severalTypesAndStrings(i int, b bool, ss ...string) {
	fmt.Println(i, b, ss)
}

func main() {
	// Println can take 1 argument
	fmt.Println(1)

	// or 5 argument
	fmt.Println(1, 2, 3, 4, 5)

	letters := []string{"a"}

	// "append" can take two arguments
	letters = append(letters, "b")

	// or 4
	letters = append(letters, "c", "d", "e")

	// twoInts accepts 2 integers
	twoInts(1, 2)
	// twoInts(1) // not enough arguments in call to twoInts have(number) want(int, int)
	// twoInts(1, 2, 3) // too many arguments in call to twoInts have(number, number, number) want(int, int)

	// severalInts accepts 1 integer
	severalInts(1) // numbers will hold the slice [1]
	// or 2 integers
	severalInts(1, 2) // numbers will hold the slice [1, 2]
	// or 3 integers
	severalInts(1, 2, 3) // numbers will hold the slice [1, 2, 3]
	// or 0 integer
	severalInts() // numbers will hold an empty slice []

	// severalStrings accepts any number of strings
	severalStrings("a", "b", "c", "d", "e")

	// severalTypesAndStrings accepts an integer, a boolean and any number of types and
	severalTypesAndStrings(7, true, "a", "b", "c", "d", "e")

	// Function caller can omit variadic arguments
	severalTypesAndStrings(1, false)
	// but non-variadic arguments are always required

	// severalTypesAndStrings(1)
	// not enough arguments in call to severalTypesAndStrings
	// have (number)
	// want (int, bool, ...string)
}

// Only the last parameter in a function definition can be variadic
// func mix(numbers ...int, flag int) {} // can only use ... with final parameter in list compiler [MisplacedDotDotDot]
