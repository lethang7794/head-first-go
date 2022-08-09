package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe(
		"0.0.0.0:8080", // Accept requests only from our own machine on port 8080.
		nil,            // The requests will be handled using functions set up via HandleFunc.
	)
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!") // Convert the string to a []byte ...
	_, err := writer.Write(message)  // ... because ResponseWriter.Write doesn't accept strings.
	if err != nil {
		log.Fatal(err)
	}
}
