package main

import (
	"log"
	"net/url"
)

//TODO: Implement Crawling Logic
func crawl(urls []URL) map[URL]([]URL) {
	//URL -> []URL
	log.Println("Starting the crawler")
	results := make(map[URL]([]URL))

	pendingSearch := urls

	rounds := 2
	for ind := rounds; ind > 0; ind-- {
		log.Println("Round ", ind)

		nextSearch := make([]URL, 0)
		for _, term := range pendingSearch {
			log.Println("Looking @ ", term)
			links := LinksFromPage(term)

			results[term] = make([]url.URL, len(links))
			copy(results[term], links)
			nextSearch = append(nextSearch, links...)
		}
		pendingSearch = make([]URL, len(nextSearch))
		copy(pendingSearch, nextSearch)
	}

	return results
}

func parallelFetch(searchURL URL, results chan<- []URL) {
	newLinks := LinksFromPage(searchURL)
	results <- newLinks
}
