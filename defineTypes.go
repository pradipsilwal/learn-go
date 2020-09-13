package main

import (
	"fmt"
	"reflect"
)

type liters float64
type gallons float64

func toGallons(l liters) gallons {
	return gallons(l * 0.264)
}

func toLiters(g gallons) liters {
	return liters(g * 3.785)
}

func definedTypes() {
	var carFuel liters
	var busFuel gallons
	carFuel = 34.5
	busFuel = 66.3
	fmt.Println("Car Fuel: ", carFuel)
	fmt.Println("Bus Fuel: ", busFuel)
	fmt.Println("carFuel Type: ", reflect.TypeOf(carFuel))
	fmt.Println("busFuel Type: ", reflect.TypeOf(busFuel))
	fmt.Printf("Convert to gallon: %0.2f\n", toGallons(carFuel))
	fmt.Printf("Convert to liters: %0.2f\n", toLiters(busFuel))
}
