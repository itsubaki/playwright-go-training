SHELL := /bin/bash

test:
	go test -v -cover $(shell go list ./... | grep -v /vendor/ | grep -v /build/ ) -coverprofile=coverage.out -covermode=atomic

install:
	go get github.com/playwright-community/playwright-go
	go run github.com/playwright-community/playwright-go/cmd/playwright install --with-deps
