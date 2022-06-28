// pass_fail reports whether a grade passing or failing
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
