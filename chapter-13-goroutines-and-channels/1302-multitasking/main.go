package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Our program makes several calls to responsiveSize, ONE AT A TIME.

// FIRST it:
// - Establishes a network connection to "example.com"
// - then waits for the site to response.
// - then prints the response size.
// ...
// ... WAIT
// ... WAIT
// ... WAIT
// ...
// THEN it:
// - Establishes a network connection to "golang.com"
// - then waits for the site to response.
// - then prints the response size.
// ...
// ... WAIT
// ... WAIT
// ... WAIT
// ...
// THEN it:
// - Establishes a network connection to "golang.com/doc"
// - then waits for the site to response.
// - then prints the response size.

// It's TOO SLOW.
// We need to SPEED it UP.
// Could we run all three calls to responseSize AT ONCE?

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
