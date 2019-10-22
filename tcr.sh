#!/usr/bin/env bash

# TCR: test && commit || revert
# codified by Kent Beck
# https://medium.com/@kentbeck_7670/test-commit-revert-870bbd756864

function test() {
    echo
    cd `git rev-parse --show-toplevel`
    go fmt ./...
    go test -v ./...
}
function commit() {
    echo
    git add . 
    git commit -m "tcr"
}
function revert() {
    echo
    git reset --hard
}

test && commit || revert
