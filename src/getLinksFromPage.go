package main

import (
	"errors"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

type (
	//URL the type that URLs are in throughout this system
	URL = url.URL
)

//TODO: Represent URLs as URLs and not Strings (reduces parsing required, keeps urls in a cleaner format)

//LinksFromPage takes a url as a string (must include the protocol)
//Returns a list of urls from the page
//Parameter and returned urls are fully formed
func LinksFromPage(urlStart URL) (links []URL) {
	res, err := http.Get(urlStart.String())
	if err != nil {
		return
	}
	defer res.Body.Close()
	tokenIterator := html.NewTokenizer(res.Body)

FetchNextToken:
	for {
		tokenType := tokenIterator.Next()

		switch {
		case tokenType == html.ErrorToken:
			return
		case tokenType == html.StartTagToken:
			token := tokenIterator.Token()

			newLinkString, linkFoundError := getAttributeValue(&token.Attr, "href")
			if linkFoundError != nil {
				continue FetchNextToken
			}

			newLinkURL, newLinkParsable := url.Parse(newLinkString)
			if newLinkParsable != nil {
				continue FetchNextToken
			}

			newLinkAbsolute := urlStart.ResolveReference(newLinkURL)
			links = append(links, *newLinkAbsolute)
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
