package main

import (
	"bufio"
	"github.com/digi-wolk/rss-feed-reader/internal/cli"
	"github.com/digi-wolk/rss-feed-reader/internal/integrations"
	"github.com/digi-wolk/rss-feed-reader/internal/output"
	"github.com/digi-wolk/rss-feed-reader/internal/rss"
	"log"
	"os"
)

func main() {
	var rssUrls []string
	args := cli.ReadCliArguments()
	if args.RSSFeedsConfigFilePath != "" {
		urls, err := readURLsFromFile(args.RSSFeedsConfigFilePath)
		if err != nil {
			log.Fatalf("Error reading URLs from file: %s", err.Error())
		}
		rssUrls = append(rssUrls, urls...)
	}
	if args.RSSFeed != "" {
		rssUrls = append(rssUrls, args.RSSFeed)
	}

	// For each URL, read the RSS feed and print the titles of the items published within the last month
	for _, url := range rssUrls {
		if args.Verbose {
			log.Printf("Reading RSS feed from: '%s'", url)
		}
		items, err := rss.ReadRSSFeed(url, args.HoursBack)
		if err != nil {
			log.Printf("Error reading RSS feed from URL '%s': %s", url, err.Error())
			continue
		}

		// Return JSON or text output based on args.OutPutType
		if args.OutPutType == "json" {
			output.PrintItemsAsJSON(items)
		} else if args.OutPutType == "text" {
			output.PrintItemsAsText(items)
		} else if args.OutPutType == "slack-comment" {
			integrations.CommentOnSlack(items, args.SlackChannel)
		} else {
			log.Printf("Error: Unknown output type '%s'", args.OutPutType)
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
