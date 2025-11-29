package main

import (
	"alojine/scooper/web"
	"alojine/scooper/writer"
)

func main() {
	domain := "www.promotions.com"

	// content := web.GetContent(domain)
	web.GetIPInfo(domain)
	writer.WriteDataToFile(domain, content)
}
