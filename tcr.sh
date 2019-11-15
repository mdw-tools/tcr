#!/usr/bin/env bash

# TCR: test && commit || revert
# codified by Kent Beck
# https://medium.com/@kentbeck_7670/test-commit-revert-870bbd756864

function test() {
    echo
    pwd
    cd `git rev-parse --show-toplevel` # navigate to top-level of git repo
    echo
    echo '-- TEST --'
    echo
    curl "http://localhost:7890/stopwatch/reset" &>/dev/null # reset timer (if running)
    make
}
function commit() {
    echo
    echo '-- COMMIT --'
    echo
    git add .
    git commit -m "tcr"
    echo
    source `dirname $0`/git.sh
    printf 'TCR commit count: %d\n' $(tcr_commit_count)
    echo
    times
    echo
    echo '-- OK --'
    return 0
}
function revert() {
    echo
    echo '-- REVERT --'
    echo
    git clean -df
    git reset --hard
    echo
    echo "Less is more." | tee >(pbcopy)
    echo
    echo '-- ERROR --'
    return 1
}

test && commit || revert
