# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint
# https://golangci-lint.run/usage/install/
lint:
	golangci-lint run ./...
.PHONY: lint

generate-mocks:
	moq -out kafkamgmt/apiv1/client/default_api_mock.go kafkamgmt/apiv1/client DefaultApi 
	moq -out kafkamgmt/apiv1/client/security_api_mock.go kafkamgmt/apiv1/client SecurityApi	
	moq -out srsmgmt/apiv1/client/registries_api_mock.go srsmgmt/apiv1/client RegistriesApi	
	moq -out kafkainstance/apiv1internal/client/default_api_mock.go kafkainstance/apiv1internal/client DefaultApi