package registrymgmt

import (
	"errors"

	registrymgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
)

// Code supplied by the API
type ServiceErrorCode string

// GetAPIError gets a strongly typed error from an error
func GetAPIError(err error) *registrymgmtclient.Error {
	var openapiError registrymgmtclient.GenericOpenAPIError

	if ok := errors.As(err, &openapiError); ok {
		errModel := openapiError.Model()

		transformedError, ok := errModel.(registrymgmtclient.Error)
		if !ok {
			return nil
		}
		return &transformedError
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
