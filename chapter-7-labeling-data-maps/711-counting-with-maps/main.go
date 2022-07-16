package main

import (
	"fmt"
	"github.com/lethang7794/head-first-go/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	fmt.Println(counts)
}
