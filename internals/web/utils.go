package web

import (
	"regexp"
	"strings"
)

func produceURL(domain string) string {
	if strings.HasPrefix(domain, "http://") || strings.HasPrefix(domain, "https://") {
		return domain
	}
	return "hhtp://" + domain
}

func StripHTMLTags(content []byte) []byte {
	re := regexp.MustCompile(`<[^>]*>`)
	return []byte(re.ReplaceAllString(string(content), ""))
}
