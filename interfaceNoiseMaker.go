package main

import "fmt"

type whistle string

func (w whistle) makeSound() {
	fmt.Println("Tweet")
}

type horn string

func (h horn) makeSound() {
	fmt.Println("Honk!")
}

type robot string

func (r robot) makeSound() {
	fmt.Println("Beep Boop")
}

func (r robot) walk() {
	fmt.Println("Powering legs")
}

type noiseMaker interface {
	makeSound()
}

func play(n noiseMaker) {
	n.makeSound()
}

func interfaceNoiseMaker() {
	var robo robot
	var horn1 horn
	var whistle1 whistle
	robo.makeSound()
	horn1.makeSound()
	whistle1.makeSound()
}
