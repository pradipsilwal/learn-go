package main

import "fmt"

func countVotesMap() {

	//Declare a map
	var myMap map[string]float64
	//Actually create the map
	myMap = make(map[string]float64)
	myMap["a"] = 1

	//map literals
	// myMap1 := map[string]float64{"a": 1, "b": 2, "c": 3}

	//Create a map and declare a variable to hold it
	ranks := make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks)
	// fmt.Println(ranks["gold"])
	// fmt.Println(ranks["silver"])
	// fmt.Println(ranks["bronze"])
	// fmt.Println("myMap1: ", myMap1)
	// val, ok := ranks["gold"]
	// fmt.Println("Ranks: ", val, " is present: ", ok)
	// delete(ranks, "gold")
	// val, ok = ranks["gold"]
	// fmt.Println("Ranks: ", val, " is present: ", ok)
	for key, value := range ranks {
		fmt.Printf("%s: %d\n", key, value)
	}
}
