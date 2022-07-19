/*
 * Apicurio Registry API [v2]
 *
 * Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 
 *
 * API version: 2.2.5.Final
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
	"fmt"
)

// ArtifactState Describes the state of an artifact or artifact version.  The following states are possible:  * ENABLED * DISABLED * DEPRECATED 
type ArtifactState string

// List of ArtifactState
const (
	ARTIFACTSTATE_ENABLED ArtifactState = "ENABLED"
	ARTIFACTSTATE_DISABLED ArtifactState = "DISABLED"
	ARTIFACTSTATE_DEPRECATED ArtifactState = "DEPRECATED"
)

var allowedArtifactStateEnumValues = []ArtifactState{
	"ENABLED",
	"DISABLED",
	"DEPRECATED",
}

func (v *ArtifactState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ArtifactState(value)
	for _, existing := range allowedArtifactStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ArtifactState", value)
}

// NewArtifactStateFromValue returns a pointer to a valid ArtifactState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewArtifactStateFromValue(v string) (*ArtifactState, error) {
	ev := ArtifactState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ArtifactState: valid values are %v", v, allowedArtifactStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ArtifactState) IsValid() bool {
	for _, existing := range allowedArtifactStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ArtifactState value
func (v ArtifactState) Ptr() *ArtifactState {
	return &v
}

type NullableArtifactState struct {
	value *ArtifactState
	isSet bool
}

func (v NullableArtifactState) Get() *ArtifactState {
	return v.value
}

func (v *NullableArtifactState) Set(val *ArtifactState) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactState) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactState(val *ArtifactState) *NullableArtifactState {
	return &NullableArtifactState{value: val, isSet: true}
}

func (v NullableArtifactState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

