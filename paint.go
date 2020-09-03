package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, length float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("Invalid width. Should be positive")
	}
	if length < 0 {
		return 0, fmt.Errorf("Invalid length. Should be positive")
	}
	amount := length * width
	return amount, nil
}

func area() {
	amount, err := paintNeeded(-4.3, 6.4)
	if err != nil {
		// fmt.Println(err)
		log.Fatal(err)
	} else {
		fmt.Printf("Amount needed %.2f", amount)
	}
}
