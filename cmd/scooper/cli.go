package main

import (
	"alojine/scooper/internals/web"
	"alojine/scooper/internals/writer"
	"flag"
	"fmt"
	"log"
	"os"
)

func runCLI() {
	domain := flag.String("domain", "", "Domain to scoop (e.g., example.com)")
	flag.Parse()

	if *domain == "" {
		fmt.Println("Usage: scooper -domain=<domain>")
		os.Exit(1)
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
