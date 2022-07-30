package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil { // There is no panic...
		return // ...so do nothing.
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	}
}

func scanDirectory(path string) { // Error return value no longer needed.
	fmt.Println(path)
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err) // Instead of returning the error value, pass it to the "panic".
	}
	for _, fileInfo := range fileInfos {
		fullPath := filepath.Join(path, fileInfo.Name())
		if fileInfo.IsDir() {
			scanDirectory(fullPath) // No more need to store or check error return value
		} else {
			fmt.Println(fullPath)
		}
	}
}

func main() {
	defer reportPanic()
	scanDirectory(os.Args[1]) // No more need to store or check error return value
}
