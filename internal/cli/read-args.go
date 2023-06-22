package cli

import (
	"flag"
	"log"
	"strings"
)

type Arguments struct {
	Verbose                bool
	RSSFeedsConfigFilePath string
	RSSFeed                string
	OutPutType             string
	SlackChannel           string
	HoursBack              int
	FilterWords            []string
	ExcludeWords           []string
	CaseInsensitive        bool
}

func ReadCliArguments() Arguments {
	verbose := flag.Bool("verbose", false, "Enable verbose mode")
	rssFeedsConfigFilePath := flag.String("rss-conf", "", "Path to RSS feeds config file")
	rssFeed := flag.String("rss-feed", "", "RSS feed URL")
	outPutType := flag.String("output", "json", "Output type: json, text, slack-comment")
	slackChannel := flag.String("slack-channel", "", "Slack channel to post comments to")
	hoursBack := flag.Int("hours-back", 0, "Number of hours to look back for new items")
	filterWords := flag.String("filter-words", "", "Comma-separated list of words to filter by")
	excludeWords := flag.String("exclude-words", "", "Comma-separated list of words to exclude")
	caseInsensitive := flag.Bool("case-insensitive", false, "Case insensitive filtering")

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
	if *hoursBack < 0 {
		log.Fatal("Number of hours should be a positive integer")
	}
	// Default: look back 24 hours
	if *hoursBack == 0 {
		*hoursBack = 24
	}
	var filterWordsArray []string
	if *filterWords != "" {
		filterWordsArray = strings.Split(*filterWords, ",")
	}
	var excludeWordsArray []string
	if *excludeWords != "" {
		excludeWordsArray = strings.Split(*excludeWords, ",")
	}

	return Arguments{
		Verbose:                *verbose,
		RSSFeedsConfigFilePath: *rssFeedsConfigFilePath,
		RSSFeed:                *rssFeed,
		OutPutType:             *outPutType,
		SlackChannel:           *slackChannel,
		HoursBack:              *hoursBack,
		FilterWords:            filterWordsArray,
		ExcludeWords:           excludeWordsArray,
		CaseInsensitive:        *caseInsensitive,
	}
}
