package web

import (
	"io"
	"log"
	"net/http"
)

const httpPrefix = "http://"

func GetContent(domain string) []byte {
	url := httpPrefix + domain

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(body))
	return body
}
