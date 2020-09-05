package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {
	votes, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(votes)
}
