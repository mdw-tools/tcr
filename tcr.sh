#!/usr/bin/env bash

# TCR: test && commit || revert
# codified by Kent Beck
# https://medium.com/@kentbeck_7670/test-commit-revert-870bbd756864

function test() {
    echo
    echo '//////////'
    echo '// TEST //'
    echo '//////////'
    echo
    curl "http://localhost:7890/stopwatch/reset" &>/dev/null # reset timer (if running)
    cd `git rev-parse --show-toplevel` # navigate to top-level of git repo
    make
}
function commit() {
    echo
    echo '////////////'
    echo '// COMMIT //'
    echo '////////////'
    echo
    git add .
    git commit -m "tcr"
    echo
    printf 'TCR commit count: %s\n' `git log --oneline origin/master..HEAD | grep tcr | wc -l | tr -d '[:space:]'`
    echo
    echo "////////"
    echo "// OK //"
    echo "////////"
    return 0
}
function revert() {
    echo
    echo '////////////'
    echo '// REVERT //'
    echo '////////////'
    echo
    echo "less is more" | tee >(pbcopy)
    echo
    git clean -df
    git reset --hard
    echo
    echo "///////////"
    echo "// ERROR //"
    echo "///////////"
    return 1
}

test && commit || revert


