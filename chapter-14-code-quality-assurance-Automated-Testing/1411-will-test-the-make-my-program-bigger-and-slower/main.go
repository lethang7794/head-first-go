package main

import "fmt"

func main() {
	fmt.Println("Isn't all this test code going to make my program BIGGER and SLOWER?")
	// NO.
	// Go test command only works with files whose names end in _test.go.
	// Other go tool command (go build, go install...) will IGNORE those _test.go file.
}
