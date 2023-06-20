package output

import (
	"fmt"
	"github.com/digi-wolk/rss-feed-reader/internal/rss"
)

func PrintItemsAsText(items []rss.Item) {
	for _, item := range items {
		fmt.Println("Title:", item.Title)
		fmt.Println("Description:", item.Description)
		fmt.Println("Link:", item.Link)
		fmt.Println("PubDate:", item.PubDate)
		fmt.Println()
	}
}
