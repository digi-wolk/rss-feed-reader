package cli

import (
	"flag"
	"log"
)

type Arguments struct {
	Verbose                bool
	RSSFeedsConfigFilePath string
	OutPutType             string
}

func ReadCliArguments() Arguments {
	verbose := flag.Bool("verbose", false, "Enable verbose mode")
	rssFeedsConfigFilePath := flag.String("rss-conf", "", "Path to RSS feeds config file")
	outPutType := flag.String("output", "json", "Output type: json, txt")

	flag.Parse()

	if *rssFeedsConfigFilePath == "" {
		log.Fatal("RSS feeds config file path is required")
	}

	return Arguments{
		Verbose:                *verbose,
		RSSFeedsConfigFilePath: *rssFeedsConfigFilePath,
		OutPutType:             *outPutType,
	}
}
