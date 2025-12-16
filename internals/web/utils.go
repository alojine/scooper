package web

import (
	"strings"
)

func produceURL(domain string) string {
	if strings.HasPrefix(domain, "http://") || strings.HasPrefix(domain, "https://") {
		return domain
	}
	return "http://" + domain
}
