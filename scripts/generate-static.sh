#!/usr/bin/env bash
## OPENAPI_FILENAME=yourapi generate.sh 

# generate an API client for a service
generate_sdk() {
    local file_name=$1
    local output_path=$2
    local package_name=$3
     
    echo "Validating OpenAPI ${file_name}"
    npx @openapitools/openapi-generator-cli validate -i "$file_name"

    echo "Generating source code based on ${file_name}"

    # remove old generated models
    rm -Rf $OUTPUT_PATH/model $OUTPUT_PATH/api
    
    npx @openapitools/openapi-generator-cli generate -g typescript-axios -i \
    "$file_name" -o "$output_path" \
    --package-name="${package_name}" \
    --additional-properties=$additional_properties \
    --ignore-file-override=.openapi-generator-ignore
}

npx @openapitools/openapi-generator-cli version-manager set 5.2.0
echo "Generating AMS SDK"
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