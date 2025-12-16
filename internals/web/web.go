package web

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
)

type IPInfo struct {
	IPv4 []string
	IPv6 []string
}

func GetHTML(domain string) ([]byte, error) {
	url := produceURL(domain)

	client := &http.Client{
		Timeout: 10 * 1e9,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to GET %s, %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func StripHTMLTags(html []byte) []byte {
	re := regexp.MustCompile(`<[^>]*>`)
	return []byte(re.ReplaceAllString(string(html), ""))
}

func GetIPInfo(domain string) (IPInfo, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return IPInfo{}, fmt.Errorf("failed to lookup domain %s: %w", domain, err)
	}

	var result IPInfo

	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			result.IPv4 = append(result.IPv4, ipv4.String())
		} else {
			result.IPv6 = append(result.IPv6, ip.String())
		}
	}

	if len(result.IPv4) == 0 && len(result.IPv6) == 0 {
		return IPInfo{}, fmt.Errorf("no IPs found for %s", domain)
	}

	return result, nil
}
