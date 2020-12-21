package main

import (
	"log"
	"net/url"
)

//TODO: Implement Crawling Logic
func crawl(urls []URL) map[URL]([]URL) {
	//URL -> []URL
	log.Println("Starting the crawler")
	pageLinks := make(map[URL]([]URL))

	pendingSearch := urls

	rounds := 5
	for ind := rounds; ind > 0; ind-- {
		log.Println("Round ", ind)
		results := make(chan parallelFetchReturnPair)
		var nextSearch []URL
		var threadCount int = 0

	SearchTerms:
		for _, term := range pendingSearch {
			_, urlAlreadyQueried := pageLinks[term]
			if urlAlreadyQueried {
				continue SearchTerms
			}
			go parallelFetch(term, results)
			threadCount++
		}
		log.Println("Amount of Links Being Searched: ", threadCount)
		for i := 0; i < threadCount; i++ {
			nextPair := <-results
			pageLinks[nextPair.url] = make([]url.URL, len(nextPair.results))
			copy(pageLinks[nextPair.url], nextPair.results)
			nextSearch = append(nextSearch, nextPair.results...)
		}
		log.Println("Amount of Links Found: ", len(nextSearch))
		pendingSearch = make([]url.URL, len(nextSearch))
		copy(pendingSearch, nextSearch)
	}

	return pageLinks
}

type parallelFetchReturnPair struct {
	url     URL
	results []URL
}

func parallelFetch(searchURL URL, results chan<- parallelFetchReturnPair) {
	newLinks := LinksFromPage(searchURL)
	output := parallelFetchReturnPair{searchURL, newLinks}
	results <- output
}
