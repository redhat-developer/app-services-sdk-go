#!/usr/bin/env bash

npx @openapitools/openapi-generator-cli version-manager set 5.2.0

echo "Generating internal AMS SDK"
additional_properties="generateInterfaces=true,enumClassPrefix=true,withSeparateModelsAndApi=true,modelPackage=model,apiPackage=api,"

OPENAPI_FILENAME=".openapi/ams.json"
PATCH_FILE=".openapi/ams.patch" 
PACKAGE_NAME="accountmgmtclient"
OUTPUT_PATH="accountmgmt/apiv1/client"

patch $OPENAPI_FILENAME < $PATCH_FILE

npx @openapitools/openapi-generator-cli generate -g go -i \
    "$OPENAPI_FILENAME" -o "$OUTPUT_PATH" \
    --package-name="${PACKAGE_NAME}" \
    --additional-properties=$additional_properties \
    --ignore-file-override=./accountmgmt/.openapi-generator-ignore 

git checkout -- $OPENAPI_FILENAME