package main

import (
	"alojine/scooper/internals/web"
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
	cfg := ParseFlags()

	html, err := web.GetContent(cfg.Domain)
	if err != nil {
		log.Fatal(err)
	}

	results := make(map[string]any)

	if cfg.Raw {
		results["raw_html"] = string(html)
	}

	if cfg.Text {
		results["text"] = string(web.StripHTMLTags(html))
	}
}
