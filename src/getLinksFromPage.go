package main

import (
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

/*
Given a url (full url, includes protocol)
Returns the links from the page
*/
func getLinksFromPage(urlString string) (links []string) {
	urlParsed, err := url.Parse(urlString)
	if err != nil {
		return
	}

	res, err := http.Get(urlString)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()

	tokenIterator := html.NewTokenizer(res.Body)
TokenLoop:
	for {
		tokenType := tokenIterator.Next()

		switch {
		case tokenType == html.ErrorToken:
			return
		case tokenType == html.StartTagToken:
			token := tokenIterator.Token()
			val, ok := getAttributeValue(&token.Attr, "href")
			if ok {
				valsParsed, err := url.Parse(val)
				if err != nil {
					continue TokenLoop
				}
				absoluteURL := urlParsed.ResolveReference(valsParsed)
				links = append(links, absoluteURL.String())
			}
		}
	}

}

/*
Given a list of attributes, find the attribute (attributeKey) from the list of attributes (attributes)
Returns value of the attribute, or empty string
Returns a valid bit as well
*/
func getAttributeValue(attributes *[]html.Attribute, attributeKey string) (string, bool) {
	for _, attribute := range *attributes {
		if attribute.Key == attributeKey {
			return attribute.Val, true
		}
	}
	return "", false
}
