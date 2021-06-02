# #!/usr/bin/env bash

# set output path of the API client
CLIENT_ID=$1
OPENAPI_FILE_URL=$2
API_GROUP="$(node $(dirname $0)/get-mapped-config.js "$CLIENT_ID" "apiGroup")"
API_VERSION="api$(node $(dirname $0)/get-mapped-config.js "$CLIENT_ID" "apiVersion")"
OPENAPI_FILENAME=$(node $(dirname $0)/get-openapi-filename.js "$OPENAPI_FILE_URL")

# set the Go package name
PACKAGE_NAME="$(node $(dirname $0)/get-mapped-config.js "$CLIENT_ID" "apiPackageName")"
if [[ ! -v "$PACKAGE_NAME" ]]; then
    echo "No package name is set, using apiGroup as package name"
    PACKAGE_NAME="$API_GROUP"
fi
OUTPUT_PATH="$API_GROUP/$API_VERSION"

npx @openapitools/openapi-generator-cli version-manager set 5.1.1
npx @openapitools/openapi-generator-cli generate -g go -i "$OPENAPI_FILENAME" -o "$OUTPUT_PATH" --package-name "$PACKAGE_NAME" --git-user-id="redhat-developer" --git-repo-id="app-services-sdk-go/$PACKAGE_NAME" -p "generateInterfaces=true" --ignore-file-override=.openapi-generator-ignore

# generate API interface mock
mock_api_file="$OUTPUT_PATH/default_api_mock.go"

if ! command -v moq &> /dev/null
then
    go get -u github.com/matryer/moq
fi

rm -rf $mock_api_file
moq -out "$mock_api_file" "$OUTPUT_PATH" DefaultApi

OPENAPI_OUTPUT_FILENAME=$OPENAPI_FILENAME

mv $OPENAPI_FILENAME .openapi/$OPENAPI_OUTPUT_FILENAME