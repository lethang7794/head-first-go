package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
)

type GuestBook struct {
	SignatureCount int
	Signatures     []string
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")

	html, err := template.ParseFiles("view.html")
	check(err)

	guestBook := GuestBook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestBook)
	check(err)
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName) // Open a file
	if os.IsNotExist(err) {        // If an error is returned saying the file doesn't exist ...
		return nil // ... return nil instead of the slice of strings.
	}
	check(err)         // For any other kind of error, report it and exit.
	defer file.Close() // After the function exits, ensure the file is closed.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err()) // Report any scanning error and exit.
	return lines
}
