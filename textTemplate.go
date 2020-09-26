package main

import (
	"datafile"
	"html/template"
	"os"
)

func textTemplate() {
	text := "Here is my template!\n"
	templ, err := template.New("test").Parse(text)
	datafile.Check(err)
	err = templ.Execute(os.Stdout, nil)
	datafile.Check(err)
}
