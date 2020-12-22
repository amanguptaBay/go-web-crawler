package main

import (
	"bufio"
	"log"
	"net/url"
	"os"
)

func runCrawl(firstURL URL) {
	cache := map[URL]bool{}
	results := crawl(
		[]URL{firstURL},
		cache,
	)

	for parentURL, links := range results {
		log.Println(parentURL, len(links))
	}
}

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
	log.Println("Crawler")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	firstURL, ok := url.Parse(text)

	if ok != nil {
		log.Fatal("Error parsing first URL")
		return
	}

	runCrawl(*firstURL)
}
