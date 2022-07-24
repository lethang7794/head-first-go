package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := Date{
		Year:  2022,
		Month: 7,
		Day:   21,
	}
	fmt.Println(date)
}
