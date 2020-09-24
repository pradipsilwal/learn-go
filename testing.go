package main

import (
	"datafile"
	"fmt"
)

func testing() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", datafile.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", datafile.JoinWithCommas(phrases))
}
