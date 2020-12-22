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
	//http://www.ics.uci.edu/~thornton/ics45c/ProjectGuide/
	log.Println("Running go-web-crawler")

	firstURL, ok := url.Parse("http://www.ics.uci.edu/~thornton/ics45c/ProjectGuide/")

	if ok != nil {
		log.Fatal("Error parsing first URL")
		return
	}
	cache := map[URL]bool{}
	pending := []URL{*firstURL}
	for range numbersUpTo(0, 4) {
		log.Println("New Round")
		log.Println("Proccessing ", len(pending), " links")
		log.Println("...")
		results := crawl(
			pending,
			cache,
		)

		for key, links := range results {
			cache[key] = true
			for _, link := range links {
				pending = append(pending, link)
			}

		}

	}
}

func numbersUpTo(start, stop int) (out []int) {
	for i := start; i < stop; i++ {
		out = append(out, i)
	}
	return
}
