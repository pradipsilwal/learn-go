package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func guess() {
	success := false
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I have chosen random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	for x := 10; x > 0; x-- {
		//Taking input
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW")
			fmt.Println("You have ", x-1, " gueses left.")
		} else if guess > target {
			fmt.Println("Oops. Your target was HIGH")
			fmt.Println("You have ", x-1, " gueses left.")
		} else {
			success = true
			break
		}
	}

	if success == true {
		fmt.Println("Good Job! You guessed it!")
	} else {
		fmt.Println("Sorry You didn't guess my number. It was: ", target)
	}
}
