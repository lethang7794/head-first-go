package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/a", f) // y
	http.HandleFunc("/b", d) // z
	http.HandleFunc("/c", e) // x

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func d(writer http.ResponseWriter, request *http.Request) {
	write(writer, "z")
}

func e(writer http.ResponseWriter, request *http.Request) {
	write(writer, "x")
}

func f(writer http.ResponseWriter, request *http.Request) {
	write(writer, "y")
}
