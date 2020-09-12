package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func printInfo(s *subscriber) {
	fmt.Println("Name: ", s.name)
	fmt.Println("Rate: ", s.rate)
	fmt.Println("Active? ", s.active)
}

func newType() {
	subscriber1 := defaultSubscriber("Pradip Silwal")
	printInfo(subscriber1)
	fmt.Println("===========================")
	applyDiscount(subscriber1)
	printInfo(subscriber1)
}
