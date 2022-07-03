package main

import "fmt"

func main() {
	fmt.Println("About one-third:", 1.0/3.0)
	fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)
	resultString := fmt.Sprintf("Still about one-third: %0.2f\n", 1.0/3.0)
	fmt.Print(resultString)
}
