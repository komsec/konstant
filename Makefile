# Go parameters
GO = go
GOCLEAN = $(GO) clean
GOTEST = $(GO) test -v
BENCHTIME ?= 2s
COVERPROFILE = cover.out
BINARY = konstant
# Get version from the tag on current HEAD
VERSION=$(shell git describe --abbrev=0)
COMMITS_AFTER_LAST_VERSION=$(shell git rev-list  `git rev-list --tags --no-walk --max-count=1`..HEAD --count)
GIT_SHA = $(shell git rev-parse --short HEAD)
GIT_DIRTY = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE = $(shell date +%Y-%m-%dT%H:%M:%S%z)

# If there are commits after last tag(version), the release is unstable.
# Also version string will be added a label with number of commits after last release tag
# e.g 0.0.1-4 (4 commits after version 0.0.1)
ifeq ($(COMMITS_AFTER_LAST_VERSION),0)
	VERSION_STRING=$(VERSION)
	RELEASE_TYPE="Stable"
else
	VERSION_STRING=$(VERSION)-$(COMMITS_AFTER_LAST_VERSION)
	RELEASE_TYPE="Unstable"
endif

LD_FLAGS = -ldflags "-X main.BuildDate=${BUILD_DATE} \
			-X main.GitCommit=${GIT_SHA} \
			-X main.GitState=${GIT_DIRTY} \
			-X main.GitSummary=${GIT_BRANCH} \
			-X main.Version=${VERSION_STRING} \
			-X main.Release=${RELEASE_TYPE}"
GOBUILD = $(GO) build $(LD_FLAGS) -o ${BINARY} -v
all: deps coverage bench build

deps:
	$(GO) get -d .

build:
	$(info Building the binary $(BINARY))
	$(GOBUILD)

build-darwin:
	$(info Building the binary $(BINARY) for Mac OSX)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD)

build-linux:
	$(info Building the binary $(BINARY) for Linux)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD)

build-win:
	$(info Building the binary $(BINARY) for WINDOWS)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD)

install:
	$(info Installing the binary $(BINARY))
	$(GO) install -v $(LD_FLAGS)

test:
	$(info Running Unit Tests)
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY).*
	rm -f $(COVERPROFILE)

bench:
	$(info Running Benchmarks)
	$(GOTEST) ./... -benchmem -run=XX -benchtime=$(BENCHTIME) -bench=.

coverage:
	$(info Running Unit tests with coverage)
	$(GOTEST) ./... -coverprofile=$(COVERPROFILE)

cover-html: coverage
	$(info Running Unit tests with html coverage report)
	$(GO) tool cover -html=$(COVERPROFILE)
	rm -f $(COVERPROFILE)

fmt:
	$(GO) fmt ./...

.PHONY: coverage cover-html deps test bench fmt
