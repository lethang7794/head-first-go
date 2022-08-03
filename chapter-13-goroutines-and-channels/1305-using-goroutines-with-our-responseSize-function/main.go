package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// Get the sizes of several pages.
	go responseSize("https://example.com")
	go responseSize("https://golang.org")
	go responseSize("https://golang.org/doc")
	time.Sleep(5 * time.Second) // Sleep for 5 seconds.
}

func responseSize(url string) {
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
	fmt.Println(len(body))
}
