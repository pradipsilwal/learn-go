package main

import (
	"fmt"
	"log"
)

func mySlice() {

	numbers, err := getFloat("data.txt")
	sum := 0.0
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range numbers {
		sum += value
	}
	countLength := float64(len(numbers))
	avg := sum / countLength
	fmt.Printf("Amount needed %.2f: ", avg)

	// var notes []string
	// notes := []string("do", "re", "mi")
	// notes := make([]string, 7)

	// notes[0] = "do"
	// notes[1] = "re"
	// notes[2] = "mi"
	// fmt.Println(notes[0])
	// fmt.Println(notes[1])
	// fmt.Println(notes[2])

	// underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	// slice1 := underlyingArray[:3]
	// fmt.Println(slice1)
	// slice1[0] = "x"
	// fmt.Println(underlyingArray)
	// fmt.Println(slice1)
}
