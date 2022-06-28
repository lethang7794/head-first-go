package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	fileInfo, err := os.Stat("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
