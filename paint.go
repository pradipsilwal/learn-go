package main

import "fmt"

func paintNeeded(width float64, length float64) float64 {
	area := length * width
	return area
}

func area() {
	fmt.Printf("Amount of paint needed: %.1f\n", paintNeeded(3.0, 5.0))
}
