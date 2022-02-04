### Go SDK uses golang wrapper for generating go code

go get -u github.com/redhat-developer/app-services-sdk-go/internal/apigen/cmd/api-generate-all@latest

api-generate-all --repo-metadata=.config/api-client-metadata.json --generator go --templates-dir=./scripts/templates