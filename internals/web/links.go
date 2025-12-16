package web

import (
	"bytes"

	"golang.org/x/net/html"
)

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
				href := attr.Val
				*urls = append(*urls, href)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walk(c, urls)
	}
}
