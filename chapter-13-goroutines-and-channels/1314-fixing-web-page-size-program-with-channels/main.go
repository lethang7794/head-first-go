package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	sizes := make(chan int)

	// Try another way
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://golang.org/doc",
	}
	for _, url := range urls {
		go responseSize(url, sizes)
	}
	for range urls {
		fmt.Println(<-sizes)
	}
}

func responseSize(url string, myChannel chan int) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	myChannel <- len(body)
}
