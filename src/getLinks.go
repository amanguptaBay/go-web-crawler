package main

import (
	"log"
	"net/http"
)

func getFromPage(url string) {

	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	res.Header.Get("Content-Type")
}
