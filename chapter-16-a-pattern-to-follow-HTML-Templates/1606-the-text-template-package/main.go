package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	text := "Here's my template!\n"
	tmpl, err := template.New("").Parse(text)
	check(err)

	// Write the template content to the terminal
	err = tmpl.Execute(os.Stdout, nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
