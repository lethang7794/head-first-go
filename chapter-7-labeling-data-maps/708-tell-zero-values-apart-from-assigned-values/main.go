package main

import "fmt"

func status(name string) {
	grades := map[string]float64{
		"Alma":  0,
		"Rohit": 86.5,
	}
	grade := grades[name]
	if grade < 60 {
		fmt.Printf("%s is failing\n", name)
	}
}

func status2(name string) {
	grades := map[string]float64{
		"Alma":  0,
		"Rohit": 86.5,
	}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing\n", name)
	}
}

func main() {
	status("Alma")
	status("Harry") // ! Harry just hasn't had any grades logged

	counters := map[string]int{
		"a": 3,
		"b": 0,
	}

	// Accessing a map key optionally returns a second, Boolean value
	value := counters["a"]
	value, ok := counters["a"]
	fmt.Printf("value: %v, ok: %v\n", value, ok)

	// The second value is true if we access an assigned value
	value, ok = counters["a"]
	fmt.Printf("value: %v, ok: %v\n", value, ok)

	// Or false if we access an unassigned value
	value, ok = counters["c"]
	fmt.Printf("value: %v, ok: %v\n", value, ok)
	// This is "comma ok idiom"

	// Test whether a a value is present
	_, ok = counters["b"]
	_, ok = counters["c"]

	status2("Alma")
	status2("Harry")
}
