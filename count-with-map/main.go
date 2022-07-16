package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/lethang7794/head-first-go/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	var counts = make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	var names []string
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s %d\n", name, counts[name])
	}
}
