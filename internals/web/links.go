package web

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

var skipPrefixes = []string{
	"#",
	"/",
	"mailto:",
	"javascript:",
}

func ExtractLinks(htmlContent []byte) ([]string, error) {
	doc, err := html.Parse(bytes.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var urls []string

	walk(doc, &urls)

	return urls, nil
}

func walk(n *html.Node, urls *[]string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				href := strings.TrimSpace(attr.Val)
				if !shouldSkip(href) {
					*urls = append(*urls, href)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walk(c, urls)
	}
}

func shouldSkip(href string) bool {
	for _, prefix := range skipPrefixes {
		check := strings.ToLower(href)
		if strings.HasPrefix(check, prefix) {
			return true
		}
	}
	return false
}
