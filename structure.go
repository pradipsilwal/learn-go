package main

import "fmt"

func structure() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	myStruct.number = 23.42
	myStruct.word = "Pradip Silwal"
	myStruct.toggle = true
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)
}
