# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint
# https://golangci-lint.run/usage/install/
lint:
	golangci-lint run ./...
.PHONY: lint

generate-mocks:
	moq -out kafkamgmt/apiv1/client/default_api_mock.go kafkamgmt/apiv1/client DefaultApi 
	moq -out kafkamgmt/apiv1/client/security_api_mock.go kafkamgmt/apiv1/client SecurityApi	
	moq -out registrymgmt/apiv1/client/registries_api_mock.go registrymgmt/apiv1/client RegistriesApi	
	moq -out kafkainstance/apiv1internal/client/topics_api_mock.go kafkainstance/apiv1internal/client TopicsApi
	moq -out kafkainstance/apiv1internal/client/groups_api_mock.go kafkainstance/apiv1internal/client GroupsApi

generate:
	api-generate-all --repo-metadata=.config/api-client-metadata.json --generator go --templates-dir=./scripts/templates
.PHONY: install