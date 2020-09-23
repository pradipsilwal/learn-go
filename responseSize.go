package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type page struct {
	urlLink string
	size    int
}

func responseSize(url string, channel chan page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- page{urlLink: url, size: len(body)}
}

func responseMain() {
	channel := make(chan page)
	urls := []string{"https://golang.org", "https://golang.org/doc", "https://example.com"}
	for _, url := range urls {
		go responseSize(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		newPage := <-channel
		fmt.Printf("%s:  %d\n", newPage.urlLink, newPage.size)
	}
}
