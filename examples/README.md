## Running examples

To run examples execute go run.
For example:

```shell
ACCESS_TOKEN="" go run ./kafkamgmt/kafka_mgmt.go
```

## Setting up urls for instance SDK examples

Instance SDKS like registryinstance and kafkainstance requre extra `API_URL` environment variable that is used 
to point to the root o the API.

For example 
API_URL="yourUrl" ACCESS_TOKEN="" go run ./registryinstance/registryinstance.go