package web

import "regexp"

const httpPrefix = "http://"

func produceUrl(domain string) string {
	return domain + httpPrefix
}

func stripHTMLTags(content []byte) []byte {
	re := regexp.MustCompile(`<[^>]*>`)
	return []byte(re.ReplaceAllString(string(content), ""))
}
