package main

import (
	"alojine/scooper/internals/web"
	"alojine/scooper/internals/writer"
	"flag"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Domain string
	Raw    bool
	Text   bool
	Meta   bool
	Links  bool
}

func ParseFlags() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.Domain, "domain", "", "Domain or URL to scrape")
	flag.BoolVar(&cfg.Raw, "r", false, "Output raw HTML")
	flag.BoolVar(&cfg.Text, "t", false, "Output just text without html tags")
	flag.BoolVar(&cfg.Meta, "m", false, "Output metadata")
	flag.BoolVar(&cfg.Links, "l", false, "Output links")

	flag.Parse()

	if cfg.Domain == "" {
		fmt.Println("Usage: scooper -domain=<domain>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	return cfg
}

func runCLI() {
	domain := flag.String("domain", "", "Domain to scoop (e.g., example.com)")
	flag.Parse()

	if *domain == "" {
		fmt.Println("Usage: scooper -domain=<domain>")
		os.Exit(1)
	}
	ipInfo, err := web.GetIPInfo(*domain)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("IPv4: ", ipInfo.IPv4)
		if ipInfo.IPv6 != nil {
			fmt.Println("IPv6: ", ipInfo.IPv6)
		}
	}

	contenent, err := web.GetContent(*domain)
	if err != nil {
		log.Fatalf("error wile getting content: %v", err)
	}

	if err := writer.WriteDataToFile(*domain, contenent); err != nil {
		log.Fatalf("Error writing data: %v", err)
	}

	fmt.Println("Scraping complete. Data saved.")
}
