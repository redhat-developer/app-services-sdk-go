# #!/usr/bin/env bash

# set output path of the API client
API_ID="$(node $(dirname $0)/get-client-payload-id.js $CLIENT_PAYLOAD)"
API_GROUP="$(node $(dirname $0)/get-mapped-config.js "$API_ID" "apiGroup")"
API_VERSION="api$(node $(dirname $0)/get-mapped-config.js "$API_ID" "apiVersion")"
OPENAPI_FILENAME=$(node $(dirname $0)/get-openapi-filename.js "$CLIENT_PAYLOAD")

# set the Go package name
PACKAGE_NAME="$(node $(dirname $0)/get-mapped-config.js "$API_ID" "apiPackageName")"
if [[ ! -v "$PACKAGE_NAME" ]]; then
    echo "No package name is set, using apiGroup as package name"
    PACKAGE_NAME="$API_GROUP"
fi
OUTPUT_PATH="$API_GROUP/$API_VERSION"

npx @openapitools/openapi-generator-cli version-manager set 5.1.1
npx @openapitools/openapi-generator-cli generate -g go -i "$OPENAPI_FILENAME" -o "$OUTPUT_PATH" --package-name "$PACKAGE_NAME" --git-user-id="redhat-developer" --git-repo-id="app-services-sdk-go/$PACKAGE_NAME" -p "generateInterfaces=true" --ignore-file-override=.openapi-generator-ignore

# generate API interface mock
mock_api_file="$OUTPUT_PATH/default_api_mock.go"

go get -u github.com/matryer/moq

rm -rf $mock_api_file
moq -out "$mock_api_file" "$OUTPUT_PATH" DefaultApi

OPENAPI_OUTPUT_FILENAME=$(node $(dirname $0)/get-openapi-filename-override.js "$CLIENT_PAYLOAD")

if [ -z ${OPENAPI_OUTPUT_FILENAME+x} ]; then
    OPENAPI_OUTPUT_FILENAME=$OPENAPI_FILENAME
fi

mv $OPENAPI_FILENAME .openapi/$OPENAPI_OUTPUT_FILENAME