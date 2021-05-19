#!/usr/bin/env bash

# set output path of the API client
API_CLIENT_KEY="$GITHUB_REPO/$OPENAPI_FILENAME"
API_GROUP="$(node $(dirname $0)/get-mapped-config.js "$API_CLIENT_KEY" "apiGroup")"
API_VERSION="$(node $(dirname $0)/get-mapped-config.js "$API_CLIENT_KEY" "apiVersion")"
# set the Go package name
PACKAGE_NAME="$(node $(dirname $0)/get-mapped-config.js "$API_CLIENT_KEY" "packageName")"
OUTPUT_PATH="$API_GROUP/$PACKAGE_NAME/$API_VERSION"

# get the download URL of the OpenAPI spec file
OPENAPI_FILE_URL=$(node $(dirname $0)/get-raw-openapi-url.js "$GITHUB_REPO" "$BRANCH" "$OPENAPI_DIRECTORY" "$OPENAPI_FILENAME")
# download the OpenAPI file
curl -H "Authorization: token $BF2_TOKEN" "$OPENAPI_FILE_URL" -o "$OPENAPI_FILENAME"
if [ $? != 0 ]; then
  exit $?
fi

npx npx @openapitools/openapi-generator-cli version-manager set 5.1.1
npx @openapitools/openapi-generator-cli generate -g go -i "$OPENAPI_FILENAME" -o "$OUTPUT_PATH" --package-name "$PACKAGE_NAME$API_VERSION" --git-user-id="redhat-developer" --git-repo-id="app-services-sdk-go/apis/kafka" -p "isGoSubmodule=true"

rm -rf "$OPENAPI_FILENAME"