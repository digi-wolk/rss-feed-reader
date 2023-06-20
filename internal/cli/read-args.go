package cli

import (
	"flag"
	"log"
)

type Arguments struct {
	Verbose                bool
	RSSFeedsConfigFilePath string
	RSSFeed                string
	OutPutType             string
	SlackChannel           string
}

func ReadCliArguments() Arguments {
	verbose := flag.Bool("verbose", false, "Enable verbose mode")
	rssFeedsConfigFilePath := flag.String("rss-conf", "", "Path to RSS feeds config file")
	rssFeed := flag.String("rss-feed", "", "RSS feed URL")
	outPutType := flag.String("output", "json", "Output type: json, text, slack-comment")
	slackChannel := flag.String("slack-channel", "", "Slack channel to post comments to")

	flag.Parse()

	if *rssFeedsConfigFilePath == "" && *rssFeed == "" {
		log.Fatal("Either RSS feed or the RSS feeds config file path is required")
	}
	if *rssFeedsConfigFilePath != "" && *rssFeed != "" {
		log.Fatal("Either RSS feed or the RSS feeds config file path should be provided, not both!")
	}
	if *outPutType != "json" && *outPutType != "text" && *outPutType != "slack-comment" {
		log.Fatalf("Unknown output type '%s'", *outPutType)
	}
	if *outPutType == "slack-comment" && *slackChannel == "" {
		log.Fatal("Slack channel is required for output type 'slack-comment'")
	}

	return Arguments{
		Verbose:                *verbose,
		RSSFeedsConfigFilePath: *rssFeedsConfigFilePath,
		RSSFeed:                *rssFeed,
		OutPutType:             *outPutType,
		SlackChannel:           *slackChannel,
	}
}
