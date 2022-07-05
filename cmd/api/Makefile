GOOS=linux
GOARCH=amd64

.PHONY: build
.PHONY: docker

GIT_COMMIT := $(shell git log -1 --format=%h)
BUILD_DATE := $(shell date +%F_%T)
VERSION := "alpha"


build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.gitCommit=$(GIT_COMMIT) -X main.buildDate=$(BUILD_DATE) -X main.version=$(VERSION)"

docker:
	CGO_ENABLED=0 go build -ldflags "-X main.gitCommit=$(GIT_COMMIT) -X main.buildDate=$(BUILD_DATE) -X main.version=$(VERSION)" -a -installsuffix cgo -o api