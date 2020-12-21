package main

import (
	"log"
	"net/url"
)

func main() {
	//content, _ := fetchPage()
	//fmt.Println("Running from main.go")
	//fmt.Printf(string(content))
	// firstURL, ok := url.Parse("http://www.google.com")
	// if ok != nil {
	// 	log.Fatal("Error parsing first URL")
	// }
	// strings := LinksFromPage(*firstURL)
	// for link := range strings {
	// 	fmt.Println(strings[link].String())
	// }
	log.Println("Crawler")

	firstURL, ok := url.Parse("http://www.ics.uci.edu/~thornton/ics45c/ProjectGuide/")
	if ok != nil {
		log.Fatal("Error parsing first URL")
		return
	}

	results := crawl(
		[]URL{*firstURL},
	)

	for parentURL, links := range results {
		log.Println(parentURL, len(links))
	}
}
