package main

import (
	"datafile"
	"fmt"
	"log"
)

func arrays() {
	numbers, err := datafile.GetFloat("data.txt")
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
}
