package main

import (
	"log"
	"os"
	"text/template"
)

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

func main() {
	templateText := "Template start\nAction: {{.}}\nTemplate end\n"
	tmpl, err := template.New("test").Parse(templateText)
	check(err)

	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)

	err = tmpl.Execute(os.Stdout, 42)
	check(err)

	err = tmpl.Execute(os.Stdout, true)
	check(err)

	executeTemplate("Dot is: {{.}}\n", false)
	executeTemplate("Dot is: {{.}}\n", 3.14)
}
