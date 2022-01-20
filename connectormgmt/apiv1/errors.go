
package connectormgmt

// connectormgmt error codes 
type ErrorCode string
  
const (

  // Forbidden to perform this action
  ERROR_4 ErrorCode = "KAFKAS-MGMT-4"

  // Forbidden to create more instances than the maximum allowed
  ERROR_5 ErrorCode = "KAFKAS-MGMT-5"

  // An entity with the specified unique values already exists
  ERROR_6 ErrorCode = "KAFKAS-MGMT-6"

  // Resource not found
  ERROR_7 ErrorCode = "KAFKAS-MGMT-7"

  // General validation failure
  ERROR_8 ErrorCode = "KAFKAS-MGMT-8"

  // Unspecified error
  ERROR_9 ErrorCode = "KAFKAS-MGMT-9"

  // HTTP Method not implemented for this endpoint
  ERROR_10 ErrorCode = "KAFKAS-MGMT-10"

  // Account is unauthorized to perform this action
  ERROR_11 ErrorCode = "KAFKAS-MGMT-11"

  // Required terms have not been accepted
  ERROR_12 ErrorCode = "KAFKAS-MGMT-12"

  // Account authentication could not be verified
  ERROR_15 ErrorCode = "KAFKAS-MGMT-15"

  // Unable to read request body
  ERROR_17 ErrorCode = "KAFKAS-MGMT-17"

  // Bad request
  ERROR_21 ErrorCode = "KAFKAS-MGMT-21"

  // Failed to parse search query
  ERROR_23 ErrorCode = "KAFKAS-MGMT-23"

  // The maximum number of allowed kafka instances has been reached
  ERROR_24 ErrorCode = "KAFKAS-MGMT-24"

  // Resource gone
  ERROR_25 ErrorCode = "KAFKAS-MGMT-25"

  // Provider not supported
  ERROR_30 ErrorCode = "KAFKAS-MGMT-30"

  // Region not supported
  ERROR_31 ErrorCode = "KAFKAS-MGMT-31"

  // Kafka cluster name is invalid
  ERROR_32 ErrorCode = "KAFKAS-MGMT-32"

  // Minimum field length not reached
  ERROR_33 ErrorCode = "KAFKAS-MGMT-33"

  // Maximum field length has been depassed
  ERROR_34 ErrorCode = "KAFKAS-MGMT-34"

  // Only multiAZ Kafkas are supported, use multi_az=true
  ERROR_35 ErrorCode = "KAFKAS-MGMT-35"

  // Kafka cluster name is already used
  ERROR_36 ErrorCode = "KAFKAS-MGMT-36"

  // Field validation failed
  ERROR_37 ErrorCode = "KAFKAS-MGMT-37"

  // Service account name is invalid
  ERROR_38 ErrorCode = "KAFKAS-MGMT-38"

  // Service account desc is invalid
  ERROR_39 ErrorCode = "KAFKAS-MGMT-39"

  // Service account id is invalid
  ERROR_40 ErrorCode = "KAFKAS-MGMT-40"

  // Instance Type not supported
  ERROR_41 ErrorCode = "KAFKAS-MGMT-41"

  // Synchronous action is not supported, use async=true parameter
  ERROR_103 ErrorCode = "KAFKAS-MGMT-103"

  // Failed to create kafka client in the mas sso
  ERROR_106 ErrorCode = "KAFKAS-MGMT-106"

  // Failed to get kafka client secret from the mas sso
  ERROR_107 ErrorCode = "KAFKAS-MGMT-107"

  // Failed to get kafka client from the mas sso
  ERROR_108 ErrorCode = "KAFKAS-MGMT-108"

  // Failed to delete kafka client from the mas sso
  ERROR_109 ErrorCode = "KAFKAS-MGMT-109"

  // Failed to create service account
  ERROR_110 ErrorCode = "KAFKAS-MGMT-110"

  // Failed to get service account
  ERROR_111 ErrorCode = "KAFKAS-MGMT-111"

  // Failed to delete service account
  ERROR_112 ErrorCode = "KAFKAS-MGMT-112"

  // Failed to find service account
  ERROR_113 ErrorCode = "KAFKAS-MGMT-113"

  // Insufficient quota
  ERROR_120 ErrorCode = "KAFKAS-MGMT-120"

  // Failed to check quota
  ERROR_121 ErrorCode = "KAFKAS-MGMT-121"

  // Too Many requests
  ERROR_429 ErrorCode = "KAFKAS-MGMT-429"

  // An unexpected error happened, please check the log of the service for details
  ERROR_1000 ErrorCode = "KAFKAS-MGMT-1000"

)