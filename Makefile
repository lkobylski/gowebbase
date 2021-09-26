
VERSION = $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT = $(shell git rev-parse --short HEAD)
#GIT_COMMIT := $(shell git rev-parse HEAD | head -c 10 | xargs)
GO_LDFLAGS = "-X version.GitCommit=$(GITCOMMIT)"
BUILD_DATE= $(shell date '+%Y-%m-%d-%H:%M:%S')
GO_TOOLS = github.com/mattn/goveralls golang.org/x/tools/cmd/cover github.com/jteeuwen/go-bindata/...


default: test

build: default

test:
	@echo "...Running tests"
	@go vet .
	@go test -v ./...

deps: bindata
	@echo "...Getting dependencies"
	@go get -t -v ./...

install: deps
	@echo "...Build and install binary"
	@echo "version $(VERSION)"
	@go install -ldflags $(GOLDFLAGS) ./...

format:
	@echo "...Format go code"
	@go fmt ./...


.PHONY: all deps format tests install bindata