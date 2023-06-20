package output

import (
	"fmt"
	"github.com/digi-wolk/rss-feed-reader/internal/rss"
	"github.com/slack-go/slack"
	"log"
	"os"
)

func CommentOnSlack(items []rss.Item, slackChannel string) {
	slackApiToken := os.Getenv("SLACK_API_TOKEN")
	if slackApiToken == "" {
		log.Fatal("Error: SLACK_API_TOKEN environment variable is not set!")
		return
	}
	api := slack.New(slackApiToken)

	for _, item := range items {
		message := buildMessage(item)
		// Post the message to the Slack channel
		_, _, err := api.PostMessage(slackChannel, slack.MsgOptionText(message, false))
		if err != nil {
			fmt.Printf("Error posting message to Slack: %s", err.Error())
		}
	}
}

func buildMessage(item rss.Item) string {
	message := fmt.Sprintf("Title: %s\n", item.Title)
	message += fmt.Sprintf("Description: %s\n", item.Description)
	message += fmt.Sprintf("Link: %s\n", item.Link)
	message += fmt.Sprintf("PubDate: %s\n", item.PubDate)

	return message
}
