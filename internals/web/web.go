package web

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

const httpPrefix = "http://"

type IPInfo struct {
	IPv4 []string
	IPv6 []string
}

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
