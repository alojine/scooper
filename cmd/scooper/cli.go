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

	flag.BoolVar(&cfg.Raw, "r", false, "Output raw HTML")
	flag.BoolVar(&cfg.Text, "t", false, "Output just text without html tags")
	flag.BoolVar(&cfg.Meta, "m", false, "Output metadata")
	flag.BoolVar(&cfg.Links, "l", false, "Output links")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: scooper <domain>")
		flag.PrintDefaults()
		os.Exit(1)
	}
	cfg.Domain = args[0]

	if !cfg.Raw && !cfg.Text && !cfg.Meta && !cfg.Links {
		cfg.Text = true
	}

	return cfg
}

func runCLI() {
	cfg := ParseFlags()

	html, err := web.GetHTML(cfg.Domain)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Raw {
		fmt.Println(string(html))
	}

	if cfg.Text {
		text := web.StripHTMLTags(html)
		fmt.Println(string(text))
	}

	if cfg.Links {
		urls, err := web.ExtractLinks(html)
		if err != nil {
			fmt.Println(err)
		}

		for _, url := range urls {
			fmt.Println(url)
		}
	}
}
