#!/usr/bin/env bash

# TCR: test && commit || revert
# codified by Kent Beck
# https://medium.com/@kentbeck_7670/test-commit-revert-870bbd756864

function test() {
    echo

    curl "http://localhost:7890/stopwatch/reset" &>/dev/null # reset timer
    cd `git rev-parse --show-toplevel` # navigate to top-level of git repo
    git status --porcelain -u | while read x # run goimports on any files that have changed
    do
      goimports -w ${x:2} # "M blah/main.go"
    done
    go fmt ./...
    go test ./...
}
function commit() {
    echo
    git add .
    git commit -m "tcr"
    return 0
}
function revert() {
    echo
    echo "less is more" | tee >(pbcopy)
    echo
    git clean -df
    git reset --hard
}

test && commit || revert
