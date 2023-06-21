<div align="center">
<img src="assets/logo/logo-no-border.png" width="250">
<br />

[![Test status](https://github.com/digi-wolk/rss-feed-reader/actions/workflows/run-tests.yml/badge.svg?branch=master)](https://github.com/digi-wolk/rss-feed-reader/actions/workflows/run-tests.yml)
[![Code quality status](https://github.com/digi-wolk/rss-feed-reader/actions/workflows/run-code-quality.yml/badge.svg?branch=master)](https://github.com/digi-wolk/rss-feed-reader/actions/workflows/run-code-quality.yml)
[![goSec security status](https://github.com/digi-wolk/rss-feed-reader/actions/workflows/run-security-checks.yml/badge.svg?branch=master)](https://github.com/digi-wolk/rss-feed-reader/actions/workflows/run-security-checks.yml)

[![License](https://img.shields.io/github/license/digi-wolk/rss-feed-reader?color=blue&label=License&style=flat-square)](https://github.com/digi-wolk/rss-feed-reader/blob/master/LICENSE.md)
[![Last commit](https://img.shields.io/github/last-commit/digi-wolk/rss-feed-reader.svg?color=blue&style=flat-square)](https://github.com/digi-wolk/rss-feed-reader/commits/master)
[![Contributors](https://img.shields.io/github/contributors/digi-wolk/rss-feed-reader?color=blue&style=flat-square)](https://github.com/digi-wolk/rss-feed-reader/graphs/contributors)

</div>

# About
Reads RSS feeds.

# Example usage
Get one RSS feed:
```bash
go run cmd/rssread/rssread.go -rss-feed=https://aws.amazon.com/about-aws/whats-new/recent/feed/
```
Get multiple RSS feeds from a config file:
```bash
go run cmd/rssread/rssread.go -rss-conf=config/rss-feeds.txt
```
Specify how many hours back to look for new items (default is 24 hours)
```bash
go run cmd/rssread/rssread.go -rss-feed=https://aws.amazon.com/about-aws/whats-new/recent/feed/ -hours-back=1
```
Specify to return text or JSON:
```bash
go run cmd/rssread/rssread.go -rss-feed=xxxx -output=text
go run cmd/rssread/rssread.go -rss-feed=xxxx -output=json
```
Verbose logs (suitable for debugging):
```bash
go run cmd/rssread/rssread.go -rss-feed=xxxx -verbose
```
Comment on Slack:
```bash
# Slack Token should be set as environment variable
export SLACK_API_TOKEN=xxxx
# Slack channel is required if output is set to slack-comment
go run cmd/rssread/rssread.go -rss-feed=https://aws.amazon.com/about-aws/whats-new/recent/feed/ -output=slack-comment -slack-channel=xxx
```
# Check the latest version
You can get the latest version number by:
- Checkin the [CHANGELOG.md](CHANGELOG.md) file
- Running this command:
```bash
git describe --abbrev=0 --tags --match "v*"
```

# Release a new version
You can create a release by tagging the version as follows:
```bash
# Update the version before running the command
RELEASE_VERSION="v1.0.0"
git tag "${RELEASE_VERSION}" -m "Release version ${RELEASE_VERSION}"
git push origin "${RELEASE_VERSION}"
```
The GitHub Action release workflow triggers immediately after the tag is pushed upstream. It will build and publish
the docker image into DockerHub.

# Docker Hub
Docker image is available on [Docker Hub](https://hub.docker.com/r/prazian/rss-feed-reader).

You can pull the latest version by running the following command:
```bash
docker pull prazian/rss-feed-reader:latest
```

You can pull a specific version by running the following command:
```bash
docker pull prazian/rss-feed-reader:v1
```

You can run the command like this:
```bash
docker run --rm -it -e GO111MODULE=on \
  prazian/rss-feed-reader:v1 \
  -rss-feed=https://aws.amazon.com/about-aws/whats-new/recent/feed/ \
  -output=slack-comment \
  -slack-channel=xxxx \
  -hours-back=1
```

# Documentation
- [Slack](docs/slack.md)