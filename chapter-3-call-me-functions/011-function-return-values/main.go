package main

import "fmt"

func double(number float64) float64 {
	return number * 2
}

func status(grade float64) string {
	if grade < 60 {
		return "failing"
	}
	return "passing"
}

func main() {
	dozen := double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))

	fmt.Println(status(60.1))
	fmt.Println(status(59.9))
}
