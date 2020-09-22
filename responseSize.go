package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}

func responseMain() {
	responseSize("https://example.com")
	responseSize("https://google.com")
	responseSize("https://facebook.com")
}
