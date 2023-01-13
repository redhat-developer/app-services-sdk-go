#!/usr/bin/env bash

# get the download URL of the OpenAPI spec file
OPENAPI_FILE_URL=$1
OPENAPI_FILE_NAME=$(basename $OPENAPI_FILE_URL)

# Special case when downloading Service Registry Instance schema file,
# to keep the original (manually updated) file name.
if [ $OPENAPI_FILE_NAME == "openapi.json" ]; then
  OPENAPI_FILE_NAME="registry-instance.json"
fi
OPENAPI_FILE_TARGET=$(dirname $0)/../.openapi/$OPENAPI_FILE_NAME

# download the OpenAPI file
curl "$OPENAPI_FILE_URL" --output $OPENAPI_FILE_TARGET
if [ $? != 0 ]; then
  exit $?
fi
