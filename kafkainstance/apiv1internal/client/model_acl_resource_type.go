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

// AclResourceType the model 'AclResourceType'
type AclResourceType string

// List of AclResourceType
const (
	ACLRESOURCETYPE_GROUP AclResourceType = "GROUP"
	ACLRESOURCETYPE_TOPIC AclResourceType = "TOPIC"
	ACLRESOURCETYPE_CLUSTER AclResourceType = "CLUSTER"
	ACLRESOURCETYPE_TRANSACTIONAL_ID AclResourceType = "TRANSACTIONAL_ID"
)

var allowedAclResourceTypeEnumValues = []AclResourceType{
	"GROUP",
	"TOPIC",
	"CLUSTER",
	"TRANSACTIONAL_ID",
}

func (v *AclResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AclResourceType(value)
	for _, existing := range allowedAclResourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AclResourceType", value)
}

// NewAclResourceTypeFromValue returns a pointer to a valid AclResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAclResourceTypeFromValue(v string) (*AclResourceType, error) {
	ev := AclResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AclResourceType: valid values are %v", v, allowedAclResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AclResourceType) IsValid() bool {
	for _, existing := range allowedAclResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AclResourceType value
func (v AclResourceType) Ptr() *AclResourceType {
	return &v
}

type NullableAclResourceType struct {
	value *AclResourceType
	isSet bool
}

func (v NullableAclResourceType) Get() *AclResourceType {
	return v.value
}

func (v *NullableAclResourceType) Set(val *AclResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableAclResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableAclResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclResourceType(val *AclResourceType) *NullableAclResourceType {
	return &NullableAclResourceType{value: val, isSet: true}
}

func (v NullableAclResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

