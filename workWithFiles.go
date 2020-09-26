package main

import (
	"bufio"
	"datafile"
	"fmt"
	"os"
)

func readOnly() {
	file, err := os.OpenFile("aardvark.txt", os.O_RDONLY, os.FileMode(0600))
	datafile.Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	datafile.Check(scanner.Err())
}

func writeOnly() {
	file, err := os.OpenFile("aardvark.txt", os.O_WRONLY, os.FileMode(0600))
	datafile.Check(err)
	_, err = file.Write([]byte("amazing!\n"))
	datafile.Check(err)
	err = file.Close()
	datafile.Check(err)
}

func appendingFile() {
	options := os.O_WRONLY | os.O_APPEND
	file, err := os.OpenFile("aardvark.txt", options, os.FileMode(0600))
	datafile.Check(err)
	_, err = file.Write([]byte("Harry Potter\n"))
	datafile.Check(err)
	err = file.Close()
	datafile.Check(err)
}

func create() {
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("aardvark.txt", options, os.FileMode(0600))
	datafile.Check(err)
	_, err = file.Write([]byte("Harry Potter\n"))
	datafile.Check(err)
	err = file.Close()
	datafile.Check(err)
}

func workWithFiles() {
	// readOnly()
	// writeOnly()
	// appendingFile()
	create()
}
