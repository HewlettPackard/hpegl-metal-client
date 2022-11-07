#!/bin/bash
# (C) Copyright 2022 Hewlett Packard Enterprise Development LP
#
# Dependencies:
#   1. yq >= 4.0.0
#      wget https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64 -O /usr/bin/yq && chmod +x /usr/bin/yq
#
#   2. redoc-cli >= 0.13.20
#      npm i redoc-cli
#

set -euo pipefail

if [ $# -lt 1 ]; then
  echo "usage: $0 <API_Spec_file_path> [<output_file_path>]" 1>&2
  exit 1
fi

OUTPUT_FILE="indext.html"
if [ ! -z "$2" ]; then
  OUTPUT_FILE=${2}
if

API_FILE=$1

TEMP_DIR=$(mktemp -d)
trap "rm -rf ${TEMP_DIR}" exit

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"

SRC_DIR=`dirname ${API_FILE}`
SRC_MAIN_FILE=`basename ${API_FILE}`

TEMP_MAIN_FILE=${TEMP_DIR}/${SRC_MAIN_FILE}

# Copying required files to temp space.
cp -r ${SRC_DIR}/* ${TEMP_DIR}
cp -r ${SCRIPT_DIR}/html ${TEMP_DIR}

# on-the-fly updating API Spec to generate offline HTML API documentation in the desired format.
yq -i '.info.description={"$ref": "./html/docs/api_description.md"}' ${TEMP_MAIN_FILE}
yq -i '.externalDocs={"description":"API Versioning", "url":"https://www.hpe.com/us/en/software/licensing.html"}' ${TEMP_MAIN_FILE}

# generate HTML file. 
redoc-cli build -o ./html/GLPCE_BMaaS_API_Spec.html ${TEMP_MAIN_FILE}
