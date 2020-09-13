package main

import (
	"datafile"
	"fmt"
)

func importingStruct() {
	subscriber1 := datafile.Subscriber{Name: "Pradip Silwal", Rate: 5.5, Active: true}
	address1 := datafile.Address{Street: "3 Gladstone Parade", City: "Glenroy", State: "VIC", PostalCode: "3046"}
	subscriber1.Address = address1
	fmt.Println("Name: ", subscriber1.Name)
	fmt.Println("Rate: ", subscriber1.Rate)
	fmt.Println("Active? ", subscriber1.Active)
	// fmt.Println("Home Address: ", subscriber1.Address)

	//Since we are using anonymous field in the struct we can
	//access the embeded struct variable using the outer structs
	fmt.Println("City: ", subscriber1.City)
}
