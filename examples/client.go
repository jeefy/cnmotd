package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	fp.UserAgent = fmt.Sprintf("%s/%s", "kubectl", "v1.28.0")
	feed, err := fp.ParseURL("http://localhost:8080?projects=kubernetes")
	if err != nil {
		log.Printf("Error parsing feed: %s", err)
	} else {
		if len(feed.Items) > 0 {
			var messageColor color.Attribute
			color.Blue("-- Cloud Native Notices --")
			for _, item := range feed.Items {
				switch item.Content {
				case "info":
					messageColor = color.FgGreen
				case "warn":
					messageColor = color.FgYellow
				case "crit":
					messageColor = color.FgRed
				}
				color.New(messageColor).Printf("%s - %s - %s\n", strings.ToUpper(item.Content), item.Title, item.Link)
			}
			color.Blue("-- /motd.cncf.io/ --")
		}
	}
}
