package main

import (
	"alojine/scooper/web"
	"alojine/scooper/writer"
)

func main() {
	domain := "www.promotions.com"

	content := web.GetContent(domain)
	writer.WriteDataToFile(domain, content)
}
