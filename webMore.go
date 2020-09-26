package main

import (
	"datafile"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	placeholder := []byte("Signature list goes here.")
	_, err := writer.Write(placeholder)
	datafile.Check(err)
}

func webMore() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
