package web

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

type Metadata struct {
	Title       string
	Description string
	Keywords    string
}

func ExtractMetadata(htmlContent []byte) (*Metadata, error) {
	doc, err := html.Parse(bytes.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	md := &Metadata{}
	walkMetadata(doc, md)
	return md, nil
}

func walkMetadata(n *html.Node, md *Metadata) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "title":
			md.Title = extractText(n)
		case "meta":
			var name, content string
			for _, attr := range n.Attr {
				switch attr.Key {
				case "name":
					name = attr.Val
				case "content":
					content = attr.Val
				}
			}
			switch strings.ToLower(name) {
			case "description":
				md.Description = content
			case "keywords":
				md.Keywords = content
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkMetadata(c, md)
	}
}

// Recursively extract text from html node, where such node might have child nodes
func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data)
	}

	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(extractText(c))
	}
	return strings.TrimSpace(sb.String())
}
