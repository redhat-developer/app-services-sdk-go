#!/usr/bin/env bash

# set output path of the API client
API_GROUP="$(node $(dirname $0)/get-mapped-config.js "$OPENAPI_FILENAME" "apiGroup")"
API_VERSION="api$(node $(dirname $0)/get-mapped-config.js "$OPENAPI_FILENAME" "apiVersion")"
# set the Go package name
PACKAGE_NAME="$(node $(dirname $0)/get-mapped-config.js "$OPENAPI_FILENAME" "apiPackageName")"
if [[ ! -v "$PACKAGE_NAME" ]]; then
    echo "No package name is set, using apiGroup as package name"
    PACKAGE_NAME="$API_VERSION"
fi
OUTPUT_PATH="$API_GROUP/$API_VERSION"

npx @openapitools/openapi-generator-cli version-manager set 5.1.1
npx @openapitools/openapi-generator-cli generate -g go -i "$OPENAPI_FILENAME" -o "$OUTPUT_PATH" --package-name "$PACKAGE_NAME" --git-user-id="redhat-developer" --git-repo-id="app-services-sdk-go/$OUTPUT_PATH" -p "generateInterfaces=true" --ignore-file-override=.openapi-generator-ignore

go mod download
go mod tidy

# generate API interface mock
mock_api_file="$OUTPUT_PATH/default_api_mock.go"

rm -rf $mock_api_file
moq -out "$mock_api_file" "$OUTPUT_PATH" DefaultApi

mv $OPENAPI_FILENAME .openapi/$OPENAPI_FILENAME