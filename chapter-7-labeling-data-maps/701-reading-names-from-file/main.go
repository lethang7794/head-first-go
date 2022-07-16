package main

import (
	"fmt"
	"log"

	"github.com/lethang7794/head-first-go/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("lines: %v\n", lines)
}
