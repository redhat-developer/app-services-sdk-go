/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
	"fmt"
)

// AclPermissionTypeFilter the model 'AclPermissionTypeFilter'
type AclPermissionTypeFilter string

// List of AclPermissionTypeFilter
const (
	ALLOW AclPermissionTypeFilter = "ALLOW"
	DENY AclPermissionTypeFilter = "DENY"
	ANY AclPermissionTypeFilter = "ANY"
)

var allowedAclPermissionTypeFilterEnumValues = []AclPermissionTypeFilter{
	"ALLOW",
	"DENY",
	"ANY",
}

func (v *AclPermissionTypeFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AclPermissionTypeFilter(value)
	for _, existing := range allowedAclPermissionTypeFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AclPermissionTypeFilter", value)
}

// NewAclPermissionTypeFilterFromValue returns a pointer to a valid AclPermissionTypeFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAclPermissionTypeFilterFromValue(v string) (*AclPermissionTypeFilter, error) {
	ev := AclPermissionTypeFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AclPermissionTypeFilter: valid values are %v", v, allowedAclPermissionTypeFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AclPermissionTypeFilter) IsValid() bool {
	for _, existing := range allowedAclPermissionTypeFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AclPermissionTypeFilter value
func (v AclPermissionTypeFilter) Ptr() *AclPermissionTypeFilter {
	return &v
}

type NullableAclPermissionTypeFilter struct {
	value *AclPermissionTypeFilter
	isSet bool
}

func (v NullableAclPermissionTypeFilter) Get() *AclPermissionTypeFilter {
	return v.value
}

func (v *NullableAclPermissionTypeFilter) Set(val *AclPermissionTypeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAclPermissionTypeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAclPermissionTypeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclPermissionTypeFilter(val *AclPermissionTypeFilter) *NullableAclPermissionTypeFilter {
	return &NullableAclPermissionTypeFilter{value: val, isSet: true}
}

func (v NullableAclPermissionTypeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclPermissionTypeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

