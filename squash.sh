#!/usr/bin/env bash

TCR_COMMIT_COUNT=`git log --oneline origin/master..HEAD | grep tcr | wc -l | tr -d '[:space:]'`
git reset --soft HEAD~$TCR_COMMIT_COUNT
