package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks"
	replacer := strings.NewReplacer("#", "o")
	fix := replacer.Replace(broken)
	fmt.Println(fix)

}
