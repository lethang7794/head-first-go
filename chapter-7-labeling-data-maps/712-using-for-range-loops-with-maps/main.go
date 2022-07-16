package main

import "fmt"

func main() {
	grades := map[string]float64{
		"Alma":  74.2,
		"Rohit": 86.5,
		"Carl":  59.7,
	}

	for name, grade := range grades {
		fmt.Printf("%s has a grade of %0.1f\n", name, grade)
	}

	for name := range grades {
		fmt.Println(name)
	}

	for _, grade := range grades {
		fmt.Println(grade)
	}
}
