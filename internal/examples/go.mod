module github.com/redhat-developer/app-services-sdk-go/internal/examples

go 1.13

replace github.com/redhat-developer/app-services-sdk-go/kafkamgmt => ../../kafkamgmt

replace github.com/redhat-developer/app-services-sdk-go/kafkainstance => ../../kafkainstance

replace github.com/redhat-developer/app-services-sdk-go/connectormgmt => ../../connectormgmt

replace github.com/redhat-developer/app-services-sdk-go/registrymgmt => ../../registrymgmt

replace github.com/redhat-developer/app-services-sdk-go/registryinstance => ../../registryinstance

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/redhat-developer/app-services-sdk-go/connectormgmt v0.0.0
	github.com/redhat-developer/app-services-sdk-go/kafkainstance v0.0.0
	github.com/redhat-developer/app-services-sdk-go/kafkamgmt v0.0.0
	github.com/redhat-developer/app-services-sdk-go/registryinstance v0.0.0
	github.com/redhat-developer/app-services-sdk-go/registrymgmt v0.0.0
	golang.org/x/oauth2 v0.0.0-20210810183815-faf39c7919d5
	google.golang.org/appengine v1.6.7 // indirect
)
