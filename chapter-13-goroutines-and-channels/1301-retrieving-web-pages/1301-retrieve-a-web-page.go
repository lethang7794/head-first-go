package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()                // Release the network connection once the "main" function exits.
	body, err := ioutil.ReadAll(response.Body) // Read all the data in the response.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body)) // Convert the data to a string and print it
}
