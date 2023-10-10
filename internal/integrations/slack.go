package integrations

import (
	"fmt"
	"github.com/digi-wolk/rss-feed-reader/internal/rss"
	"github.com/slack-go/slack"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
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
	description := removeHTMLTags(item.Description)

	titleBlock := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Title:*\n%s", item.Title), false, false)
	descriptionBlock := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Description:*\n%s", description), false, false)
	linkBlock := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Link:*\n%s", item.Link), false, false)
	pubDateBlock := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*PubDate:*\n%s", item.PubDate), false, false)

	return []slack.Block{
		slack.NewSectionBlock(titleBlock, nil, nil),
		slack.NewSectionBlock(descriptionBlock, nil, nil),
		slack.NewSectionBlock(linkBlock, nil, nil),
		slack.NewSectionBlock(pubDateBlock, nil, nil),
	}
}

func removeHTMLTags(htmlContent string) string {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Printf("Error parsing HTML: %s", err.Error())
		return htmlContent // Return original content as fallback
	}

	textContent := ""
	var extractText func(*html.Node)
	extractText = func(n *html.Node) {
		if n.Type == html.TextNode {
			textContent += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c != nil {
				extractText(c)
			}
		}
	}

	extractText(doc)
	return strings.TrimSpace(textContent)
}
