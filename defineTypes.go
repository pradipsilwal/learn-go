package main

import (
	"fmt"
	"reflect"
)

type litersVar float64
type gallonsVar float64

func toGallons(l litersVar) gallonsVar {
	return gallonsVar(l * 0.264)
}

func toLiters(g gallonsVar) litersVar {
	return litersVar(g * 3.785)
}

func definedTypes() {
	var carFuel litersVar
	var busFuel gallonsVar
	carFuel = 34.5
	busFuel = 66.3
	fmt.Println("Car Fuel: ", carFuel)
	fmt.Println("Bus Fuel: ", busFuel)
	fmt.Println("carFuel Type: ", reflect.TypeOf(carFuel))
	fmt.Println("busFuel Type: ", reflect.TypeOf(busFuel))
	fmt.Printf("Convert to gallon: %0.2f\n", toGallons(carFuel))
	fmt.Printf("Convert to liters: %0.2f\n", toLiters(busFuel))
}
