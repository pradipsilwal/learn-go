package main

import (
	"html/template"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func textTemplate() {
	text := "Here is my template!\n"
	templ, err := template.New("test").Parse(text)
	check(err)
	err = templ.Execute(os.Stdout, nil)
	check(err)
}
