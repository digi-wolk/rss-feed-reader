package output

import (
	"encoding/json"
	"fmt"
	"github.com/digi-wolk/rss-feed-reader/internal/rss"
)

func PrintItemsAsJSON(items []rss.Item) {
	jsonData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		fmt.Printf("Error converting items to JSON: %s", err.Error())
		return
	}
	fmt.Println(string(jsonData))
}
