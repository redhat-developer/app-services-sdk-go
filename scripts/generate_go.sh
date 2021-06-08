# #!/usr/bin/env bash

# set output path of the API client
CLIENT_ID=$1
OPENAPI_FILE_URL=$2
TEMPLATES_DIR="$(dirname $0)/templates"
API_GROUP="$(node $(dirname $0)/get-mapped-config.js "$CLIENT_ID" "apiGroup")"
API_VERSION="api$(node $(dirname $0)/get-mapped-config.js "$CLIENT_ID" "apiVersion")"
OPENAPI_FILENAME=$(node $(dirname $0)/get-openapi-filename.js "$OPENAPI_FILE_URL")

# set the Go package name
PACKAGE_NAME="$(node $(dirname $0)/get-mapped-config.js "$CLIENT_ID" "apiPackageName")"
if [[ -z "$PACKAGE_NAME" ]] || [[ "$PACKAGE_NAME" == "undefined" ]]; then
    echo "No package name is set, using apiGroup as package name"
    PACKAGE_NAME="${API_GROUP}"
fi
PACKAGE_NAME="${PACKAGE_NAME}client"
OUTPUT_PATH="$API_GROUP/$API_VERSION/client"

npx @openapitools/openapi-generator-cli version-manager set 5.1.1
npx @openapitools/openapi-generator-cli generate -g go -i "$OPENAPI_FILENAME" -o "$OUTPUT_PATH" -t "$TEMPLATES_DIR" --package-name "$PACKAGE_NAME" --git-user-id="redhat-developer" --git-repo-id="app-services-sdk-go/$PACKAGE_NAME" -p "generateInterfaces=true" --ignore-file-override=.openapi-generator-ignore

# generate API interface mock
mock_api_file="$OUTPUT_PATH/default_api_mock.go"

go get -u github.com/matryer/moq
rm -rf $mock_api_file
moq -out "$mock_api_file" "$OUTPUT_PATH" DefaultApi

go mod tidy

OPENAPI_OUTPUT_FILENAME=$OPENAPI_FILENAME

mv $OPENAPI_FILENAME .openapi/$OPENAPI_OUTPUT_FILENAME