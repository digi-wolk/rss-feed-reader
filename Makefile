GIT_TAG?= $(shell git describe --always --tags)
BUILD_FLAGS := "-w -s -X 'main.Version=$(GIT_TAG)' -X 'main.GitTag=$(GIT_TAG)' -X 'main.BuildDate=$(BUILD_DATE)'"
BIN=rssread
CGO_ENABLED = 0
GO_VERSION = 1.20

build-linux-amd64:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -ldflags=$(BUILD_FLAGS) -o build/$(BIN) ./cmd/rssread/

build-linux-arm64:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=arm64 go build -ldflags=$(BUILD_FLAGS) -o build/$(BIN) ./cmd/rssread/

# Update all dependencies to their latest versions
update-deps:
	go get -u ./...
	go mod tidy

# Update a specific dependency to its latest version
# Usage: make update-pkg PKG=github.com/example/package
update-pkg:
	@if [ -z "$(PKG)" ]; then \
		echo "Error: PKG is not set. Usage: make update-pkg PKG=github.com/example/package"; \
		exit 1; \
	fi
	go get -u $(PKG)
	go mod tidy

# List all dependencies and their versions
list-deps:
	go list -m all
