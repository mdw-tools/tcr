#!/usr/bin/make -f

test: install
	go test -count=1 -short $(ARGS) ./...

install:
	go install github.com/mdwhatcott/tcr.sh/cmd/...
