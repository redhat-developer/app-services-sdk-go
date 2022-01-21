package kafkamgmt

import (
	"errors"
	"fmt"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

var (
	NotFoundByIDErr         error
	NotFoundByNameErr       error
	IllegalSearchValueError error
	InvalidNameErr          error
)

// Code supplied by the API
type ServiceErrorCode string

// GetAPIError gets a strongly typed error from an error
func GetAPIError(err error) *kafkamgmtclient.Error {
	var openapiError kafkamgmtclient.GenericOpenAPIError

	if ok := errors.As(err, &openapiError); ok {
		errModel := openapiError.Model()

		kafkaMgmtError, ok := errModel.(kafkamgmtclient.Error)
		if !ok {
			return nil
		}
		return &kafkaMgmtError
	}

	return nil
}

// IsAPIError returns true if the error contains the errCode
func IsAPIError(err error, code ServiceErrorCode) bool {
	mappedErr := GetAPIError(err)
	if mappedErr == nil {
		return false
	}

	return mappedErr.GetCode() == string(code)
}

func NotFoundByIDError(id string) error {
	NotFoundByIDErr = fmt.Errorf(`Kafka instance with ID "%v" not found`, id)
	return NotFoundByIDErr
}

func NotFoundByNameError(name string) error {
	NotFoundByNameErr = fmt.Errorf(`Kafka instance "%v" not found`, name)
	return NotFoundByNameErr
}

func InvalidSearchValueError(v string) error {
	IllegalSearchValueError = fmt.Errorf(`
	illegal search value "%v", search input must satisfy the following conditions:

  - must be of 1 or more characters
  - must only consist of alphanumeric characters, '-', '_' and '%%'
	`, v)

	return IllegalSearchValueError
}

func InvalidNameError(v string) error {
	InvalidNameErr = fmt.Errorf(`invalid Kafka instance name "%v". Valid names must satisfy the following conditions:

  - must be between 1 and 32 characters
  - must only consist of lower case, alphanumeric characters and '-'
  - must start with an alphabetic character
  - must end with an alphanumeric character
	`, v)
	return InvalidNameErr
}
