package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
		return err
	}
	for _, fileInfo := range fileInfos {
		fullPath := filepath.Join(path, fileInfo.Name())
		if fileInfo.IsDir() {
			err := scanDirectory(fullPath)
			// If an error occurs in one of the recursive scanDirectory calls,
			// that error has to be returned up the ENTIRE CHAIN until it reaches the main function!
			if err != nil {
				fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
				return err
			}
		} else {
			fmt.Println(fullPath)
		}
	}
	return nil
}

func main() {
	err := scanDirectory(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
