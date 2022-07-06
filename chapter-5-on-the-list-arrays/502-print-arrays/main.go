package main

import "fmt"

func main() {
	var notes [3]string = [3]string{"do", "re", "mi"}
	var primes [5]int = [5]int{2, 3, 5, 7, 11}

	fmt.Println(notes)
	fmt.Println(primes)
	
	fmt.Printf("%#v\n", notes)
	fmt.Printf("%#v\n", primes)
}
