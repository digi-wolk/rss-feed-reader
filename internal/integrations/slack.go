package integrations

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
		messageBlocks := buildMessageBlocks(item)
		// Post the message to the Slack channel
		_, _, err := api.PostMessage(slackChannel, slack.MsgOptionBlocks(messageBlocks...))
		if err != nil {
			fmt.Printf("Error posting message to Slack: %s", err.Error())
		}
	}
}

func buildMessageBlocks(item rss.Item) []slack.Block {
	titleBlock := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Title:*\n%s", item.Title), false, false)
	descriptionBlock := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Description:*\n%s", item.Description), false, false)
	linkBlock := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Link:*\n%s", item.Link), false, false)
	pubDateBlock := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*PubDate:*\n%s", item.PubDate), false, false)

	return []slack.Block{
		slack.NewSectionBlock(titleBlock, nil, nil),
		slack.NewSectionBlock(descriptionBlock, nil, nil),
		slack.NewSectionBlock(linkBlock, nil, nil),
		slack.NewSectionBlock(pubDateBlock, nil, nil),
	}
}
