## Prerequisites:

Run the following command(s) to ensure proper functionality of the program. 

```shell
go mod tidy
```


## Running examples

To run examples execute go run.
For example:

```shell
ACCESS_TOKEN=`rhoas authtoken` go run ./examples/kafkamgmt/kafka_mgmt.go
```

## Setting up urls for instance SDK examples

Instance SDKS like registryinstance and kafkainstance requre extra `API_URL` environment variable that is used to point to the root of the API.

For example 
API_URL="yourUrl" ACCESS_TOKEN=`rhoas authtoken` go run ./examples/registryinstance/registryinstance.go


