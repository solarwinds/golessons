.PHONY: build test

excluding_vendor := $(shell go list ./... | grep -v /vendor/)

default: build

build:
	go build -i

test:
	go test -v $(excluding_vendor)
