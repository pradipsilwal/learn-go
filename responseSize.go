package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Getting", url)
	fmt.Println(len(body))
}

func responseMain() {
	go responseSize("https://example.com")
	go responseSize("https://google.com")
	go responseSize("https://facebook.com")
	time.Sleep(5 * time.Second)
}
