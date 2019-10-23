#!/usr/bin/env bash

TCR_COMMIT_COUNT=`git log --oneline origin/master..HEAD | grep tcr | wc -l | tr -d '[:space:]'`
printf "Squashing [%s] commits into single, staged change set...\n" $TCR_COMMIT_COUNT
git reset --soft HEAD~$TCR_COMMIT_COUNT
echo
git status