package kafkamgmt

// KafkaManagementErrorCode represents error code returned by Kafka Management API
type KafkaManagementErrorCode string

const (
	// Forbidden to perform this action
	ERROR_4 KafkaManagementErrorCode = "KAFKAS-MGMT-4"

	// Forbidden to create more instances than the maximum allowed
	ERROR_5 KafkaManagementErrorCode = "KAFKAS-MGMT-5"

	// An entity with the specified unique values already exists
	ERROR_6 KafkaManagementErrorCode = "KAFKAS-MGMT-6"

	// Resource not found
	ERROR_7 KafkaManagementErrorCode = "KAFKAS-MGMT-7"

	// General validation failure
	ERROR_8 KafkaManagementErrorCode = "KAFKAS-MGMT-8"

	// Unspecified error
	ERROR_9 KafkaManagementErrorCode = "KAFKAS-MGMT-9"

	// HTTP Method not implemented for this endpoint
	ERROR_10 KafkaManagementErrorCode = "KAFKAS-MGMT-10"

	// Account is unauthorized to perform this action
	ERROR_11 KafkaManagementErrorCode = "KAFKAS-MGMT-11"

	// Required terms have not been accepted
	ERROR_12 KafkaManagementErrorCode = "KAFKAS-MGMT-12"

	// Account authentication could not be verified
	ERROR_15 KafkaManagementErrorCode = "KAFKAS-MGMT-15"

	// Unable to read request body
	ERROR_17 KafkaManagementErrorCode = "KAFKAS-MGMT-17"

	// Bad request
	ERROR_21 KafkaManagementErrorCode = "KAFKAS-MGMT-21"

	// Failed to parse search query
	ERROR_23 KafkaManagementErrorCode = "KAFKAS-MGMT-23"

	// The maximum number of allowed kafka instances has been reached
	ERROR_24 KafkaManagementErrorCode = "KAFKAS-MGMT-24"

	// Resource gone
	ERROR_25 KafkaManagementErrorCode = "KAFKAS-MGMT-25"

	// Provider not supported
	ERROR_30 KafkaManagementErrorCode = "KAFKAS-MGMT-30"

	// Region not supported
	ERROR_31 KafkaManagementErrorCode = "KAFKAS-MGMT-31"

	// Kafka cluster name is invalid
	ERROR_32 KafkaManagementErrorCode = "KAFKAS-MGMT-32"

	// Minimum field length not reached
	ERROR_33 KafkaManagementErrorCode = "KAFKAS-MGMT-33"

	// Maximum field length has been depassed
	ERROR_34 KafkaManagementErrorCode = "KAFKAS-MGMT-34"

	// Only multiAZ Kafkas are supported, use multi_az=true
	ERROR_35 KafkaManagementErrorCode = "KAFKAS-MGMT-35"

	// Kafka cluster name is already used
	ERROR_36 KafkaManagementErrorCode = "KAFKAS-MGMT-36"

	// Field validation failed
	ERROR_37 KafkaManagementErrorCode = "KAFKAS-MGMT-37"

	// Service account name is invalid
	ERROR_38 KafkaManagementErrorCode = "KAFKAS-MGMT-38"

	// Service account desc is invalid
	ERROR_39 KafkaManagementErrorCode = "KAFKAS-MGMT-39"

	// Service account id is invalid
	ERROR_40 KafkaManagementErrorCode = "KAFKAS-MGMT-40"

	// Instance Type not supported
	ERROR_41 KafkaManagementErrorCode = "KAFKAS-MGMT-41"

	// Synchronous action is not supported, use async=true parameter
	ERROR_103 KafkaManagementErrorCode = "KAFKAS-MGMT-103"

	// Failed to create kafka client in the mas sso
	ERROR_106 KafkaManagementErrorCode = "KAFKAS-MGMT-106"

	// Failed to get kafka client secret from the mas sso
	ERROR_107 KafkaManagementErrorCode = "KAFKAS-MGMT-107"

	// Failed to get kafka client from the mas sso
	ERROR_108 KafkaManagementErrorCode = "KAFKAS-MGMT-108"

	// Failed to delete kafka client from the mas sso
	ERROR_109 KafkaManagementErrorCode = "KAFKAS-MGMT-109"

	// Failed to create service account
	ERROR_110 KafkaManagementErrorCode = "KAFKAS-MGMT-110"

	// Failed to get service account
	ERROR_111 KafkaManagementErrorCode = "KAFKAS-MGMT-111"

	// Failed to delete service account
	ERROR_112 KafkaManagementErrorCode = "KAFKAS-MGMT-112"

	// Failed to find service account
	ERROR_113 KafkaManagementErrorCode = "KAFKAS-MGMT-113"

	// Insufficient quota
	ERROR_120 KafkaManagementErrorCode = "KAFKAS-MGMT-120"

	// Failed to check quota
	ERROR_121 KafkaManagementErrorCode = "KAFKAS-MGMT-121"

	// Too Many requests
	ERROR_429 KafkaManagementErrorCode = "KAFKAS-MGMT-429"

	// An unexpected error happened, please check the log of the service for details
	ERROR_1000 KafkaManagementErrorCode = "KAFKAS-MGMT-1000"
)
