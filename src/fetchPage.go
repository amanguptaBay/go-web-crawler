package main

import (
	"net/http"
	"log"
	"io/ioutil"
)

func fetchPage() ([]byte, error){
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return robots, err
}