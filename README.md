# About
Reads RSS feeds from config files.

# Example usage
Get one RSS feed:
```bash
go run cmd/read-rss.go -rss-feed=https://aws.amazon.com/about-aws/whats-new/recent/feed/
```
Get multiple RSS feeds from a config file:
```bash
go run cmd/read-rss.go -rss-conf=config/rss-feeds.txt
```
Specify to return text or JSON:
```bash
go run cmd/read-rss.go -rss-feed=xxxx -output=text
go run cmd/read-rss.go -rss-feed=xxxx -output=json
```
Verbose logs (suitable for debugging):
```bash
go run cmd/read-rss.go -rss-feed=xxxx -verbose
```
Comment on Slack:
```bash
# Slack Token should be set as environment variable
export SLACK_API_TOKEN=xxxx
# Slack channel is required if output is set to slack-comment
go run cmd/read-rss.go -rss-conf=../config/rss-feeds.txt -output=slack-comment -slack-channel=xxxx
```

# Dockerfile
TBD