package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Get the sizes of several pages.
	responseSize("https://example.com")
	responseSize("https://golang.org")
	responseSize("https://golang.org/doc")
}

// Move the code that gets the page to a separate function.
func responseSize(url string) { // Take the URL as a parameter.
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
