# RHOAS SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/redhat-developer/app-services-sdk-go.svg)](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-go)

Go packages and API clients for RHOAS.

## Prequisites

- [Go 1.13](https://golang.org/doc/go1.13) or above

## Installation

Install the RHOAS SDK with `go get`:

```shell
$ go get github.com/redhat-developer/app-services-sdk-go
```

```go
import "github.com/redhat-developer/app-services-sdk-go"
```

## Supported APIs

> NOTE: Some of these APIs are under development and may sometimes cause backwards-incompatible changes.


| API                       | Status | Package |
|---------------------------|--------|---------|
| KafkaManagement           | beta |  [`github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1)      |
| ServiceRegistryManagement | alpha  |  [`github.com/redhat-developer/app-services-sdk-go/serviceregistrymgmt/apiv1`](https://pkg.go.dev/github.com/redhat-developer/app-services-sdk-go/serviceregistrymgmt/apiv1)       |

**Alpha**: the API is still being developed and may have backwards-incompatible changes. It is not recommended for production use.

**Beta**: the API is mostly complete and stable, but still has outstanding features to be addressed. There may be minor backwards-incomptable changes.

**Stable**: the API is ready for production use.

## Contributing

Contributions are welcome. See [CONTRIBUTING](CONTRIBUTING.md) for details.