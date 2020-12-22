package main

import (
	"time"
)

func crawl(urls []URL, cache map[URL]bool) (results map[URL]([]URL)) {
	results = make(map[URL]([]URL))

	workerResultChannel := make(chan parallelFetchReturnPair)
	workerCount := 0
	for _, link := range urls {
		_, inCache := cache[link]
		if inCache {
			continue
		}
		go parallelFetch(link, workerResultChannel)
		workerCount++
	}
FetchLoop:
	for i := 0; i < workerCount; i++ {
		select {
		case result := <-workerResultChannel:
			results[result.url] = make([]URL, len(result.results))
			results[result.url] = result.results
		case <-time.After(10 * time.Second):
			break FetchLoop
		}

	}
	return
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
