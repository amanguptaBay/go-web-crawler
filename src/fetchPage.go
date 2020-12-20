package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func fetchPage() ([]byte, error) {
	res, err := http.Get("http://www.google.com/robots.txt")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	return robots, err
}
