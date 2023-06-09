package main

import (
	"fmt"
	"log"

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
			color.Blue("-- Cloud Native Notices --")
			for _, item := range feed.Items {
				fmt.Println("- ", color.YellowString(item.Title), " - ", color.GreenString(item.Link))
			}
			color.Blue("-- /motd.cncf.io/ --")
		}
	}
}
