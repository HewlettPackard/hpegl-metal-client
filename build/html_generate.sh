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

if [ $# -ne 1 ]; then
  echo "usage: $0 <API_Spec_file_path>" 1>&2
  exit 1
fi

API_FILE=$1

TEMP_API_FILE=$(mktemp --suffix=".yaml")
trap "rm -f ${TEMP_API_FILE}" exit

# copy to temp file before editing.
cp ${API_FILE} ${TEMP_API_FILE}

# on-the-fly updating API Spec to generate offline HTML API documentation in the desired format.
yq -i '.info.description={"$ref": "./html/docs/api_description.md"}' ${TEMP_API_FILE}
yq -i '.externalDocs={"description":"API Versioning", "url":"https://www.hpe.com/us/en/software/licensing.html"}' ${TEMP_API_FILE}

# generate HTML file. 
redoc-cli build -o ./html/GLPCE_BMaaS_API_Spec.html ${TEMP_API_FILE}
