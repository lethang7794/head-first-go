package main

import "fmt"

func main() {
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))

	// Add an element to a slice with append
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))

	// Add some elements to a slice
	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))

	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	s5 := append(s4, "s5", "s5")

	fmt.Println(s1, s2, s3, s4, s5)

	s2[0] = "22"
	s3[0] = "33"
	s4[0] = "44"
	s5[0] = "55"
	fmt.Println(s1, s2, s3, s4, s5) // ! s3 and s4 has the same underlying

	// Just assign the return value of append to the same slice variable passed to append
	s1 = append(s1, "s2", "s2")
	s1 = append(s1, "s3", "s3")
	s1[0] = "33"
	fmt.Println(s1)

	s1 = append(s1, "s4", "s4")
	s1[0] = "44"
	fmt.Println(s1)
}
