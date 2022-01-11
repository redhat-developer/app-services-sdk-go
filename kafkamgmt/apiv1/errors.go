
package kafkamgmt

// KafkaManagementErrorCode represents error code returned by Kafka Management API
type ServiceErrorCode string
  
const (
    // Forbidden to perform this action
  ERROR_4 ServiceErrorCode = "KAFKAS-MGMT-4"

  // Forbidden to create more instances than the maximum allowed
  ERROR_5 ServiceErrorCode = "KAFKAS-MGMT-5"

  // An entity with the specified unique values already exists
  ERROR_6 ServiceErrorCode = "KAFKAS-MGMT-6"

  // Resource not found
  ERROR_7 ServiceErrorCode = "KAFKAS-MGMT-7"

  // General validation failure
  ERROR_8 ServiceErrorCode = "KAFKAS-MGMT-8"

  // Unspecified error
  ERROR_9 ServiceErrorCode = "KAFKAS-MGMT-9"

  // HTTP Method not implemented for this endpoint
  ERROR_10 ServiceErrorCode = "KAFKAS-MGMT-10"

  // Account is unauthorized to perform this action
  ERROR_11 ServiceErrorCode = "KAFKAS-MGMT-11"

  // Required terms have not been accepted
  ERROR_12 ServiceErrorCode = "KAFKAS-MGMT-12"

  // Account authentication could not be verified
  ERROR_15 ServiceErrorCode = "KAFKAS-MGMT-15"

  // Unable to read request body
  ERROR_17 ServiceErrorCode = "KAFKAS-MGMT-17"

  // Bad request
  ERROR_21 ServiceErrorCode = "KAFKAS-MGMT-21"

  // Failed to parse search query
  ERROR_23 ServiceErrorCode = "KAFKAS-MGMT-23"

  // The maximum number of allowed kafka instances has been reached
  ERROR_24 ServiceErrorCode = "KAFKAS-MGMT-24"

  // Resource gone
  ERROR_25 ServiceErrorCode = "KAFKAS-MGMT-25"

  // Provider not supported
  ERROR_30 ServiceErrorCode = "KAFKAS-MGMT-30"

  // Region not supported
  ERROR_31 ServiceErrorCode = "KAFKAS-MGMT-31"

  // Kafka cluster name is invalid
  ERROR_32 ServiceErrorCode = "KAFKAS-MGMT-32"

  // Minimum field length not reached
  ERROR_33 ServiceErrorCode = "KAFKAS-MGMT-33"

  // Maximum field length has been depassed
  ERROR_34 ServiceErrorCode = "KAFKAS-MGMT-34"

  // Only multiAZ Kafkas are supported, use multi_az=true
  ERROR_35 ServiceErrorCode = "KAFKAS-MGMT-35"

  // Kafka cluster name is already used
  ERROR_36 ServiceErrorCode = "KAFKAS-MGMT-36"

  // Field validation failed
  ERROR_37 ServiceErrorCode = "KAFKAS-MGMT-37"

  // Service account name is invalid
  ERROR_38 ServiceErrorCode = "KAFKAS-MGMT-38"

  // Service account desc is invalid
  ERROR_39 ServiceErrorCode = "KAFKAS-MGMT-39"

  // Service account id is invalid
  ERROR_40 ServiceErrorCode = "KAFKAS-MGMT-40"

  // Instance Type not supported
  ERROR_41 ServiceErrorCode = "KAFKAS-MGMT-41"

  // Synchronous action is not supported, use async=true parameter
  ERROR_103 ServiceErrorCode = "KAFKAS-MGMT-103"

  // Failed to create kafka client in the mas sso
  ERROR_106 ServiceErrorCode = "KAFKAS-MGMT-106"

  // Failed to get kafka client secret from the mas sso
  ERROR_107 ServiceErrorCode = "KAFKAS-MGMT-107"

  // Failed to get kafka client from the mas sso
  ERROR_108 ServiceErrorCode = "KAFKAS-MGMT-108"

  // Failed to delete kafka client from the mas sso
  ERROR_109 ServiceErrorCode = "KAFKAS-MGMT-109"

  // Failed to create service account
  ERROR_110 ServiceErrorCode = "KAFKAS-MGMT-110"

  // Failed to get service account
  ERROR_111 ServiceErrorCode = "KAFKAS-MGMT-111"

  // Failed to delete service account
  ERROR_112 ServiceErrorCode = "KAFKAS-MGMT-112"

  // Failed to find service account
  ERROR_113 ServiceErrorCode = "KAFKAS-MGMT-113"

  // Insufficient quota
  ERROR_120 ServiceErrorCode = "KAFKAS-MGMT-120"

  // Failed to check quota
  ERROR_121 ServiceErrorCode = "KAFKAS-MGMT-121"

  // Too Many requests
  ERROR_429 ServiceErrorCode = "KAFKAS-MGMT-429"

  // An unexpected error happened, please check the log of the service for details
  ERROR_1000 ServiceErrorCode = "KAFKAS-MGMT-1000"

)