#! /bin/bash
# (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

if [ $# -ne 1 ]; then
  echo "usage: $0 <directory>" 1>&2
  exit 1
fi

GENERATED_DIR=$1

diff="$(git status -s ${GENERATED_DIR})";
if [[ ! -z $diff ]]; then
    echo "$diff";
    exit 1;
fi
