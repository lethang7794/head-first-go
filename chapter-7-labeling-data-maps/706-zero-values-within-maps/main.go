package main

import "fmt"

func main() {
	numbers := make(map[string]int)
	numbers["I've have been assigned"] = 12

	fmt.Printf("%#v\n", numbers["I've have been assigned"])
	fmt.Printf("%#v\n", numbers["I've haven't been assigned"])

	strings := make(map[string]string)
	strings["I've have been assigned"] = "hi"

	fmt.Printf("%#v\n", strings["I've have been assigned"])
	fmt.Printf("%#v\n", strings["I've haven't been assigned"])

	counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["b"], counters["c"])
}
