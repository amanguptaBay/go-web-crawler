package main

import (
	"errors"
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

			newLinkString, linkFoundError := getAttributeValue(&token.Attr, "href")
			if linkFoundError != nil {
				continue TokenLoop
			}

			newLinkURL, newLinkParsable := url.Parse(newLinkString)
			if newLinkParsable != nil {
				continue TokenLoop
			}

			newLinkAbsolute := urlParsed.ResolveReference(newLinkURL)
			links = append(links, newLinkAbsolute.String())
		}
	}
}

/*
Given a list of attributes, find the attribute (attributeKey) from the list of attributes (attributes)
Returns value of the attribute, or empty string
Returns a valid bit as well
*/
func getAttributeValue(attributes *[]html.Attribute, attributeKey string) (string, error) {
	for _, attribute := range *attributes {
		if attribute.Key == attributeKey {
			return attribute.Val, nil
		}
	}
	return "", errors.New("getAttributeValue(): No attribute " + attributeKey + " found")
}
