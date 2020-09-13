package main

import "fmt"

type liters float64
type milliliters float64
type gallons float64

func (l liters) toGallons() gallons {
	return gallons(l * 0.264)
}

func (m milliliters) toGallons() gallons {
	return gallons(m * 0.000264)
}

func (g gallons) toLiters() liters {
	return liters(g * 3.785)
}

func defineTypeMethod() {
	var carFuel liters
	var busFuel gallons
	carFuel = 2
	busFuel = 66.3
	water := milliliters(500)

	carFuelGallons := carFuel.toGallons()
	busFuelLiters := busFuel.toLiters()
	waterGallons := water.toGallons()

	fmt.Println("Car Fuel liters: ", carFuel)
	fmt.Println("Car Fuel gallons: ", carFuelGallons)
	fmt.Println("Bus Fuel liters: ", busFuelLiters)
	fmt.Println("Bus Fuel gallons: ", busFuel)
	fmt.Println("Water in millileters: ", water)
	fmt.Println("Water in gallons: ", waterGallons)
}
