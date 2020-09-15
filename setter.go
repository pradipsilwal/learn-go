package main

import (
	"datafile"
	"fmt"
	"log"
)

func setter() {
	objDate := datafile.Event{}
	err := objDate.SetTitle("This is a very very long title written for demo purposes.")
	if err != nil {
		log.Fatal(err)
	}
	err = objDate.SetYear(1994)
	if err != nil {
		log.Fatal(err)
	}
	err = objDate.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	err = objDate.SetDay(6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(objDate.Title(), " : ")
	fmt.Print(objDate.Year(), "/")
	fmt.Print(objDate.Month(), "/")
	fmt.Println(objDate.Day())

}
