package main

import "fmt"

func myVariadic(a int, b int, c ...string) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
