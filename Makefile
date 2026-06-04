.PHONY: build test test-cover lint vulncheck clean all

GO      ?= go
BINARY  := calc

build:
	$(GO) build -ldflags="-s -w" -o $(BINARY) ./cmd/

test:
	$(GO) test -race -coverprofile=coverage.out -covermode=atomic ./...

test-cover: test
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

lint:
	golangci-lint run ./...

vulncheck:
	govulncheck ./...

clean:
	rm -f $(BINARY) coverage.out coverage.html

all: lint test build
