package main

import (
	"time"
	"fmt"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
}

