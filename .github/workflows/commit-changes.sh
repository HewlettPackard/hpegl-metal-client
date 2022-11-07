#!/bin/bash -x
# (C) Copyright 2022 Hewlett Packard Enterprise Development LP
#

set -eu

WATCH_DIR=${WATCH_DIR:-.}

if [[ -n `git status -s ${WATCH_DIR}` ]]
then
    git add ${WATCH_DIR}
    git commit -m "${GIT_COMMIT_MESSAGE:-commit changes}"
else
    echo "${WATCH_DIR} is already up-to-date"
    exit 0
fi

REMOTE="${REMOTE:-https://x-access-token:${GITHUB_TOKEN}@github.com/${REPOSITORY}}"

MAX_RETRIES=${MAX_RETRIES:-3}
ATTEMPT=0
while true
do
    if git push "${REMOTE}"
    then
        break
    fi

    ATTEMPT=$((ATTEMPT + 1))
    if [[ ${ATTEMPT} -ge ${MAX_RETRIES} ]]
    then
        echo "Failed to successfully sync with remote after ${MAX_RETRIES} attempts"
        exit 1
    fi

    echo "Try ${ATTEMPT}: failed to push, waiting 2 seconds before fetching latest and retrying..."

    sleep 2

    git pull --rebase "${REMOTE}"
done

echo "Successfully pushed latest changes to remote"