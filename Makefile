#!/usr/bin/make -f

VERSION := $(shell git describe)

test: install
	go test -count=1 -short $(ARGS) ./...

install:
	go install -ldflags="-X 'main.Version=$(VERSION)'" github.com/mdwhatcott/tcr/cmd/...
