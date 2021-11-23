ERRORS_FOLDER=$(dirname $0)/../.errors/

cd $ERRORS_FOLDER

echo "Fetching error specifications from backends"

curl https://api.openshift.com/api/kafkas_mgmt/v1/errors --output errors_$CLIENT_ID.json

## https://github.com/bf2fc6cc711aee1a0c2a/srs-fleet-manager/issues/175
# curl https://api.openshift.com/api/serviceregistry_mgmt/v1/errors --output errors_$CLIENT_ID.json

## Not enabled yet
# curl https://api.openshift.com/api/connector_mgmt/v1/errors --output errors_$CLIENT_ID.json

echo "Generating golang code"

node generate-files.js

echo "Successfully generated all code"