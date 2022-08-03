package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct { // Declare a struct and the fields we need.
	URL  string
	Size int
}

func responseSize(url string, myChannel chan Page) { // The channel pass to responseSize will carry Pages, not ints.
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

	// Send back a Page with both the current URl and the page size.
	myChannel <- Page{
		URL:  url,
		Size: len(body),
	}
}

func main() {
	pages := make(chan Page)

	// Try another way
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://golang.org/doc",
	}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for range urls {
		page := <-pages
		fmt.Println(page.URL, page.Size)
	}
}
