# Contribution Guide

## Autogeneration of SDKs

1. All SDKs are autogenerated using OpenAPI Generator project.
2. Bot accounts creates PR with the new versions of the SDK that needs be reviewed by the team
3. Once PR is merged developers are responsible for creating tagged release 


## Release process

1. Check what SDK changed
2. Create tag
3. Generate changelog and put it into release description

## Contributing to App Services SDK

## Generating an API client locally

You may want to generate an SDK locally. This can be done in two steps.

1. Fetch the OpenAPI specification file.

```shell
# Set the arguments
remote_openapi_url="https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/main/openapi/kas-fleet-manager.yaml"

# List of client IDs can be found in `.config/api-client-metadata.json`
client_id="kafka-mgmt/v1"

./scripts/fetch_api.sh $remote_openapi_url
```

2. Generate the OpenAPI client

```shell
./scripts/generate_go.sh $client_id $remote_openapi_url
```