package web

import (
	"bytes"

	"golang.org/x/net/html"
)

type Metadata struct {
	Title       string
	Description string
	Keywords    string
	Author      string
}

func ExtractMetadata(htmlContent []byte) (*Metadata, error) {
	doc, err := html.Parse(bytes.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}
}
