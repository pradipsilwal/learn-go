package main

import "fmt"

func arrays() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	for _, value := range numbers {
		sum += value
	}
	sampleCount := float64(len(numbers))
	avg := sum / sampleCount
	fmt.Printf("%.2f", avg)
}
