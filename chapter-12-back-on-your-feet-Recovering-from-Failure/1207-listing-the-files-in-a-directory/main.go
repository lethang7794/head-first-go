package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileInfos, err := ioutil.ReadDir("my_directory")
	if err != nil {
		return
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			fmt.Println("Directory:", fileInfo.Name())
		} else {
			fmt.Println("File:", fileInfo.Name())
		}

	}
}
