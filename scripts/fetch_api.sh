#!/usr/bin/env bash

# get the download URL of the OpenAPI spec file
OPENAPI_FILE_URL=$1
OPENAPI_FILE_NAME=$(node $(dirname $0)/get-openapi-filename.js "$OPENAPI_FILE_URL")

# download the OpenAPI file
curl -H "Authorization: token $BF2_TOKEN" "$OPENAPI_FILE_URL" -o "$OPENAPI_FILE_NAME"

if [ $? != 0 ]; then
  exit $?
fi
