#!/usr/bin/env bash

# get the download URL of the OpenAPI spec file
OPENAPI_FILE_URL=$(node $(dirname $0)/get-raw-openapi-url.js "$CLIENT_PAYLOAD")
OPENAPI_FILE_NAME=$(node $(dirname $0)/get-openapi-filename.js "$CLIENT_PAYLOAD")

# download the OpenAPI file
curl -H "Authorization: token $BF2_TOKEN" "$OPENAPI_FILE_URL" -o "$OPENAPI_FILE_NAME"
if [ $? != 0 ]; then
  exit $?
fi
