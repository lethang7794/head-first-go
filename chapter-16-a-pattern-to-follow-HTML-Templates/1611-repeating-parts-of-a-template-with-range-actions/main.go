package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//												---------------- <- This portion will be repeated once for each element in the slice.
	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
	//									 ___________         -----  _______
	executeTemplate(templateText, []string{"do", "re", "mi"})

	//						            -------- <- This portion will be repeated once for each element in the slice.
	templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	//						 ___________ -----  _______
	executeTemplate(templateText, []float64{1.11, 2.22, 3.33})
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
