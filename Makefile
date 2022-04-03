.PHONY: all
all: build test run

.PHONY: build
build:
	go build -o ./bin/slowly ./cmd/main

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run ./cmd/main 123.45 USD BTC
