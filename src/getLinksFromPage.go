package main

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

/*
Given a url (full url, includes protocol)
Returns the links from the page
*/
func getLinksFromPage(url string) (links []string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()

	tokenIterator := html.NewTokenizer(res.Body)

	for {
		tokenType := tokenIterator.Next()

		switch {
		case tokenType == html.ErrorToken:
			return
		case tokenType == html.StartTagToken:
			token := tokenIterator.Token()
			val, ok := getAttributeValue(token.Attr, "href")
			if ok {
				links = append(links, val)
			}
		}
	}

}

/*
Given a list of attributes, find the attribute (attributeKey) from the list of attributes (attributes)
Returns value of the attribute, or empty string
Returns a valid bit as well
*/
func getAttributeValue(attributes []html.Attribute, attributeKey string) (string, bool) {
	for _, attribute := range attributes {
		if attribute.Key == attributeKey {
			return attribute.Val, true
		}
	}
	return "", false
}
