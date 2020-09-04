package main

import (
	"bufio"
	"os"
	"strconv"
)

func readFile(fileName string) ([3]float64, error) {
	var numbers [3]float64
	i := 0
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}

	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, err
	}
	return numbers, err
}
