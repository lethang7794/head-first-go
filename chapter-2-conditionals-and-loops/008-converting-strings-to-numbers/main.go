// pass_fail reports whether a grade passing or failing
package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)


func main() {
	fmt.Print("Enter a grade: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	
	var status string
	if grade >= 60 {
		status = "passing"
		fmt.Println(status)
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
