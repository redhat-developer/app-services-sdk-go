/*
 * Red Hat Openshift SmartEvents Fleet Manager
 *
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * API version: 0.0.3
 * Contact: openbridge-dev@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsclient

import (
	"encoding/json"
	"fmt"
)

// ManagedResourceStatus the model 'ManagedResourceStatus'
type ManagedResourceStatus string

// List of ManagedResourceStatus
const (
	MANAGEDRESOURCESTATUS_ACCEPTED ManagedResourceStatus = "accepted"
	MANAGEDRESOURCESTATUS_PREPARING ManagedResourceStatus = "preparing"
	MANAGEDRESOURCESTATUS_PROVISIONING ManagedResourceStatus = "provisioning"
	MANAGEDRESOURCESTATUS_READY ManagedResourceStatus = "ready"
	MANAGEDRESOURCESTATUS_DEPROVISION ManagedResourceStatus = "deprovision"
	MANAGEDRESOURCESTATUS_DELETING ManagedResourceStatus = "deleting"
	MANAGEDRESOURCESTATUS_DELETED ManagedResourceStatus = "deleted"
	MANAGEDRESOURCESTATUS_FAILED ManagedResourceStatus = "failed"
)

var allowedManagedResourceStatusEnumValues = []ManagedResourceStatus{
	"accepted",
	"preparing",
	"provisioning",
	"ready",
	"deprovision",
	"deleting",
	"deleted",
	"failed",
}

func (v *ManagedResourceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ManagedResourceStatus(value)
	for _, existing := range allowedManagedResourceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ManagedResourceStatus", value)
}

// NewManagedResourceStatusFromValue returns a pointer to a valid ManagedResourceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewManagedResourceStatusFromValue(v string) (*ManagedResourceStatus, error) {
	ev := ManagedResourceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ManagedResourceStatus: valid values are %v", v, allowedManagedResourceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ManagedResourceStatus) IsValid() bool {
	for _, existing := range allowedManagedResourceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ManagedResourceStatus value
func (v ManagedResourceStatus) Ptr() *ManagedResourceStatus {
	return &v
}

type NullableManagedResourceStatus struct {
	value *ManagedResourceStatus
	isSet bool
}

func (v NullableManagedResourceStatus) Get() *ManagedResourceStatus {
	return v.value
}

func (v *NullableManagedResourceStatus) Set(val *ManagedResourceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedResourceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedResourceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedResourceStatus(val *ManagedResourceStatus) *NullableManagedResourceStatus {
	return &NullableManagedResourceStatus{value: val, isSet: true}
}

func (v NullableManagedResourceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedResourceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

