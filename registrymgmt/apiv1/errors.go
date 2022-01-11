
package registrymgmt

// KafkaManagementErrorCode represents error code returned by Kafka Management API
type ServiceErrorCode string
  
const (
    // Unspecified error
  ERROR_1 ServiceErrorCode = "SRS-MGMT-1"

  // Registry with id='?' not found
  ERROR_2 ServiceErrorCode = "SRS-MGMT-2"

  // Bad date or time format
  ERROR_3 ServiceErrorCode = "SRS-MGMT-3"

  // Invalid request content. Make sure the request conforms to the given JSON schema
  ERROR_4 ServiceErrorCode = "SRS-MGMT-4"

  // Bad request format - invalid JSON
  ERROR_5 ServiceErrorCode = "SRS-MGMT-5"

  // Required terms have not been accepted for account id='?'
  ERROR_6 ServiceErrorCode = "SRS-MGMT-6"

  // The maximum number of allowed Registry instances has been reached
  ERROR_7 ServiceErrorCode = "SRS-MGMT-7"

  // Error type with id='?' not found
  ERROR_8 ServiceErrorCode = "SRS-MGMT-8"

  // Data conflict. Make sure a Registry with the given name does not already exist
  ERROR_9 ServiceErrorCode = "SRS-MGMT-9"

  // Bad request format - unsupported media type
  ERROR_10 ServiceErrorCode = "SRS-MGMT-10"

)