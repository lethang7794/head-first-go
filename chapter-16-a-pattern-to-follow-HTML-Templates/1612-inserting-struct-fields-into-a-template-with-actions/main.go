package main

import (
	"log"
	"os"
	"text/template"
)

type Part struct {
	Name  string
	Count int
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func main() {
	// An action with dot followed by a field name will insert that field's value in the template.
	//						 .Name <- Insert the value of the Part's Name field.
	templateText := "Name: {{.Name}} - Count: {{.Count}}\n"
	//						 					.Count <- Insert the value of the Part's Count field.

	executeTemplate(templateText, Part{"Fuses", 5})
	executeTemplate(templateText, Part{"Cables", 2})

	templateText = "Name: {{.Name}}{{if .Active}} - Rate: ${{.Rate}}\n{{end}}"
	// 											 _____________________
	// This portion will be output only if the Subscriber's Active field value is true
	subscriber := Subscriber{
		Name:   "Aman Singh",
		Rate:   4.99,
		Active: true,
	}
	executeTemplate(templateText, subscriber)

	subscriber = Subscriber{
		Name:   "Joy Carr",
		Rate:   5.99,
		Active: false,
	}
	executeTemplate(templateText, subscriber)
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
