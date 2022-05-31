## Prerequisites:

Run the following command(s) to ensure proper functionality of the program. 

```shell
go mod tidy
```


## Running examples

To run examples execute go run, using ACCESS_TOKEN along with API_URL where necessary.

### For example:

### For kafkamgmt: 

```shell
ACCESS_TOKEN=`rhoas authtoken` go run ./examples/kafkamgmt/kafka_mgmt.go
```

### For connectormgmt:

```shell
ACCESS_TOKEN=`rhoas authtoken` go run ./examples/connectormgmt/connector_mgmt.go
```

### For srsmgmt:

```shell
ACCESS_TOKEN=`rhoas authtoken` go run ./examples/srsmgmt/srs_mgmt.go
```

## Setting up urls for instance SDK examples

Instance SDKS like registryinstance and kafkainstance requre extra `API_URL` environment variable that is used to point to the root of the API.

This can be found using the following command:
```shell
rhoas service-registry use <registry name>

rhoas service-registry describe
```
and the intended URL can be found under: "registryURL"



### For example

### For registryinstance
```shell
API_URL="yourUrl" ACCESS_TOKEN=`rhoas authtoken` go run ./examples/registryinstance/registryinstance.go
```

### For kafkainstance
```shell
API_URL="yourUrl" ACCESS_TOKEN=`rhoas authtoken` go run ./examples/kafkainstance/kafka_instance.go
```

#### Troubleshooting:

One of the common issues one can encounter during use is incomplete commands due to lack of resources (eg: Service Registry artifacts)

This can be resolved by following the steps mentioned in the following file which explains RHOAS "how-to's" as well as a good documentation for first-time users.

>Red Hat OpenShift Service Registry documentation: 
https://access.redhat.com/documentation/en-us/red_hat_openshift_service_registry/1/guide/1630486d-c8d0-45d0-b2de-7fdb52c78499#proc-creating-service-registry-account_getting-started-rhoas-service-registry
