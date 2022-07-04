# RHOAS SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/redhat-developer/app-services-sdk-go.svg)](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-go)

Go packages and API clients for Red Had OpenShift Application Services (RHOAS) 

## Prequisites

- [Go 1.15](https://golang.org/doc/go1.15) or above

## Installation

Install the RHOAS SDK with `go get`:

```shell
$ go get github.com/redhat-developer/app-services-sdk-go
```

Import:

```go
import "github.com/redhat-developer/app-services-sdk-go"
```

## Management SDKs

> NOTE: Some of these APIs are under development and may sometimes cause backwards-incompatible changes.


| API                       | Status | Package                                                                                                                                                         |
| :------------------------ | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| KafkaManagement           | beta   | [`github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1)         |
| Service Accounts Management  | alpha   | [`github.com/redhat-developer/app-services-sdk-go/serviceaccounts/apiv1`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-go/serviceaccounts/apiv1)         |
| ServiceRegistryManagement | alpha  | [`github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1)   |
| ConnectorManagement       | alpha  | [`github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1) |

 
 ## Instances SDKs

| API              | Status | Package                                                                                                                                                                               |
| ---------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| KafkaInstance    | beta   | [`github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal)       |
| RegistryInstance | beta   | [`github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal) |


## Documentation

[Documentation](./docs)

## Examples

[Examples](./examples)

## Contributing

Contributions are welcome. See [CONTRIBUTING](CONTRIBUTING.md) for details.
