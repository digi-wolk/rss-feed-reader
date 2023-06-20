# About
Reads RSS feeds from config files.

# Example usage
Print output:
```bash
go run read-rss.go -rss-conf=../config/rss-feeds.txt -output=txt
go run read-rss.go -rss-conf=../config/rss-feeds.txt -output=json
```
Comment on Slack
```bash
# Slack Token should be set as environment variable
export SLACK_API_TOKEN=xxxx
# Slack channel is required if output is set to slack-comment
go run read-rss.go -rss-conf=../config/rss-feeds.txt -output=slack-comment -slack-channel=xxxx
```

# Dockerfile
TBD