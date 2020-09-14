package main

import (
	"datafile"
	"fmt"
	"log"
)

func setter() {
	objDate := datafile.Date{}
	err := objDate.SetYear(1994)
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
	fmt.Print(objDate.Year(), "/")
	fmt.Print(objDate.Month(), "/")
	fmt.Println(objDate.Day())

}
