package main

import (
	"bufio"
	"github.com/digi-wolk/rss-feed-reader/internal/cli"
	"github.com/digi-wolk/rss-feed-reader/internal/output"
	"github.com/digi-wolk/rss-feed-reader/internal/rss"
	"log"
	"os"
)

func main() {
	args := cli.ReadCliArguments()
	filePath := args.RSSFeedsConfigFilePath
	urls, err := readURLsFromFile(filePath)
	if err != nil {
		log.Fatalf("Error reading URLs from file: %s", err.Error())
	}

	// For each URL, read the RSS feed and print the titles of the items published within the last month
	for _, url := range urls {
		items, err := rss.ReadRSSFeed(url)
		if err != nil {
			log.Printf("Error reading RSS feed from URL '%s': %s", url, err.Error())
			continue
		}

		// Return JSON or text output based on args.OutPutType
		if args.OutPutType == "json" {
			output.PrintItemsAsJSON(items)
		} else {
			output.PrintItemsAsText(items)
		}
	}
}

// Helper function to read URLs from a file
func readURLsFromFile(filePath string) ([]string, error) {
	var urls []string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}
