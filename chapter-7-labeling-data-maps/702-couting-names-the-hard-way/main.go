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

	var names []string
	var count []int
	for _, line := range lines {
		found := false
		for i, name := range names {
			if name == line {
				found = true
				count[i]++
			}
		}
		if !found {
			names = append(names, line)
			count = append(count, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%v %v\n", name, count[i])
	}
}
