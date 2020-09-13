package main

import (
	"datafile"
	"fmt"
)

func importingStruct() {
	subscriber1 := datafile.Subscriber{Name: "Pradip Silwal", Rate: 5.5, Active: true}
	fmt.Println("Name: ", subscriber1.Name)
	fmt.Println("Rate: ", subscriber1.Rate)
	fmt.Println("Active? ", subscriber1.Active)
}
