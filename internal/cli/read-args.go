package cli

import (
	"flag"
	"log"
)

type Arguments struct {
	Verbose                bool
	RSSFeedsConfigFilePath string
	OutPutType             string
	SlackChannel           string
}

func ReadCliArguments() Arguments {
	verbose := flag.Bool("verbose", false, "Enable verbose mode")
	rssFeedsConfigFilePath := flag.String("rss-conf", "", "Path to RSS feeds config file")
	outPutType := flag.String("output", "json", "Output type: json, text, slack-comment")
	slackChannel := flag.String("slack-channel", "", "Slack channel to post comments to")

	flag.Parse()

	if *rssFeedsConfigFilePath == "" {
		log.Fatal("RSS feeds config file path is required")
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
		OutPutType:             *outPutType,
		SlackChannel:           *slackChannel,
	}
}
