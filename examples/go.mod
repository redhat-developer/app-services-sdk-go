module github.com/redhat-developer/app-services-sdk-go/internal/examples

go 1.15

replace github.com/redhat-developer/app-services-sdk-go/kafkamgmt => ../kafkamgmt

replace github.com/redhat-developer/app-services-sdk-go/kafkainstance => ../kafkainstance

replace github.com/redhat-developer/app-services-sdk-go/connectormgmt => ../connectormgmt

replace github.com/redhat-developer/app-services-sdk-go/registrymgmt => ../registrymgmt

replace github.com/redhat-developer/app-services-sdk-go/registryinstance => ../registryinstance

replace github.com/redhat-developer/app-services-sdk-go/auth => ../auth

require (
	github.com/redhat-developer/app-services-sdk-go/auth v0.0.0-00010101000000-000000000000
	github.com/redhat-developer/app-services-sdk-go/connectormgmt v0.0.0
	github.com/redhat-developer/app-services-sdk-go/kafkainstance v0.0.0
	github.com/redhat-developer/app-services-sdk-go/kafkamgmt v0.0.0
	github.com/redhat-developer/app-services-sdk-go/registryinstance v0.0.0
	github.com/redhat-developer/app-services-sdk-go/registrymgmt v0.0.0
)
