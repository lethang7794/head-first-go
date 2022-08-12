package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//                     ____________ <- This portion will appear only if the dot value is true.
	tmpl := "start {{if .}}Dot is true!{{end}} finish\n"
	executeTemplate(tmpl, true)
	executeTemplate(tmpl, false)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}
