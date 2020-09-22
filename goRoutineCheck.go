package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func goRoutineCheck() {
	go a()
	go b()
	fmt.Println("End Main()")
	time.Sleep(time.Second)
}
