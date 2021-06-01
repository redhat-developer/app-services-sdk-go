#!/usr/bin/env bash

echo $CLIENT_PAYLOAD

# get the download URL of the OpenAPI spec file
OPENAPI_FILE_URL=$(echo $CLIENT_PAYLOAD | jq -r .download_url)
OPENAPI_FILE_NAME=$(node $(dirname $0)/get-openapi-filename.js "$CLIENT_PAYLOAD")

# download the OpenAPI file
curl -H "Authorization: token $BF2_TOKEN" "$OPENAPI_FILE_URL" -o "$OPENAPI_FILE_NAME"

if [ $? != 0 ]; then
  exit $?
fi
