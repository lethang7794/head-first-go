package main

import (
	"fmt"
	"github.com/lethang7794/headfirstgo/date"
)

func main() {
	date1 := date.Date{
		Year:  2022,
		Month: 14, // Invalid
		Day:   50, // Invalid
	}
	fmt.Println(date1)

	date2 := date.Date{
		Year:  -1, // Invalid
		Month: -2, // Invalid
		Day:   -3, // Invalid
	}
	fmt.Println(date2)

}
