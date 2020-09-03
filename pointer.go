package main

import "fmt"

//Return pointer
func createPointer() *int {
	myInt := 14
	return &myInt
}

func pointers() {
	var myInt int

	//Declaring pointer of type int
	var myIntPointer *int

	//Holds the address of myInt
	myIntPointer = &myInt

	myInt = 12
	fmt.Println(myInt)
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)

	//Change the value at the location pointed by myIntPointer
	*myIntPointer = 5
	fmt.Println(myInt)

	myIntPointer = createPointer()
	fmt.Println(*myIntPointer)
}
