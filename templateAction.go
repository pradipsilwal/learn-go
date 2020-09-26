package main

import (
	"datafile"
	"html/template"
	"os"
)

type Part struct {
	Name  string
	Count int
}

func templateAction1() {
	templateText := "Template Start \nAction: {{.}}\nTemplate end\n"
	tmpl, err := template.New("test").Parse(templateText)
	datafile.Check(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	datafile.Check(err)
	err = tmpl.Execute(os.Stdout, 42)
	datafile.Check(err)
	err = tmpl.Execute(os.Stdout, true)
	datafile.Check(err)
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	datafile.Check(err)
	err = tmpl.Execute(os.Stdout, data)
}

func templateActionMain() {
	// templateAction1()
	// executeTemplate("Dot is: {{.}}\n", "ABC")
	// executeTemplate("Dot is: {{.}}\n", 123.5)

	// If action in template
	// executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", true)
	// executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", false)

	// looping with range action
	// templateText := "Before loop: {{.}}\n{{range .}}In Loop: {{.}}\n{{end}}After loop: {{.}}\n"
	// executeTemplate(templateText, []string{"do", "re", "mi"})
	// Another method for looping
	// templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	// executeTemplate(templateText, []float64{1.25, 0.99, 27})

	// Inserting struct field into a template action
	templateText := "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateText, Part{Name: "Pradip", Count: 1})
	executeTemplate(templateText, Part{Name: "Prakash", Count: 2})
}
