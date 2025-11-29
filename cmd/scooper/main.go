package main

import (
	"alojine/scooper/internals/web"
	"alojine/scooper/internals/writer"
)

func main() {
	domain := "www.promotions.com"
	content := web.GetContent(domain)
	web.GetIPInfo(domain)
	writer.WriteDataToFile(domain, content)
}
