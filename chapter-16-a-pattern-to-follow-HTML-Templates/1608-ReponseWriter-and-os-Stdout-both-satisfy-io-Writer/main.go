package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	_, err := os.Stdout.Write([]byte("hello\n")) // Write data to the terminal.
	check(err)

	tmpl, err := template.New("test").Parse("Here's my template!\n")
	check(err)

	// Write of the template text ...
	err = tmpl.Execute(os.Stdout, nil) // ... to the terminal,...
	check(err)

	http.HandleFunc(
		"/guestbook",
		func(writer http.ResponseWriter, request *http.Request) {
			err := tmpl.Execute(writer, nil) // ... or to the HTTP reponse.
			check(err)
		},
	)
	err = http.ListenAndServe("localhost:8080", nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
