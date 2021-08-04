/*
 * Apicurio Registry API [v2]
 *
 * Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 
 *
 * API version: 2.1.0-SNAPSHOT
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
)

// SearchedArtifact Models a single artifact from the result set returned when searching for artifacts.
type SearchedArtifact struct {

	// The ID of a single artifact.
	Id string `json:"id"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	CreatedOn string `json:"createdOn"`

	CreatedBy string `json:"createdBy"`

	Type ArtifactType `json:"type"`

	Labels *[]string `json:"labels,omitempty"`

	State ArtifactState `json:"state"`

	ModifiedOn *string `json:"modifiedOn,omitempty"`

	ModifiedBy *string `json:"modifiedBy,omitempty"`

	// An ID of a single artifact group.
	GroupId *string `json:"groupId,omitempty"`

}

// NewSearchedArtifact instantiates a new SearchedArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchedArtifact(id string, createdOn string, createdBy string, type_ ArtifactType, state ArtifactState) *SearchedArtifact {
	this := SearchedArtifact{}
	this.Id = id
	this.CreatedOn = createdOn
	this.CreatedBy = createdBy
	this.Type = type_
	this.State = state
	return &this
}

// NewSearchedArtifactWithDefaults instantiates a new SearchedArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchedArtifactWithDefaults() *SearchedArtifact {
	this := SearchedArtifact{}












	return &this
}


// GetId returns the Id field value
func (o *SearchedArtifact) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SearchedArtifact) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SearchedArtifact) SetId(v string) {
	o.Id = v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchedArtifact) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchedArtifact) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchedArtifact) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchedArtifact) SetName(v string) {
	o.Name = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SearchedArtifact) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchedArtifact) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SearchedArtifact) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SearchedArtifact) SetDescription(v string) {
	o.Description = &v
}


// GetCreatedOn returns the CreatedOn field value
func (o *SearchedArtifact) GetCreatedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *SearchedArtifact) GetCreatedOnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *SearchedArtifact) SetCreatedOn(v string) {
	o.CreatedOn = v
}


// GetCreatedBy returns the CreatedBy field value
func (o *SearchedArtifact) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SearchedArtifact) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SearchedArtifact) SetCreatedBy(v string) {
	o.CreatedBy = v
}


// GetType returns the Type field value
func (o *SearchedArtifact) GetType() ArtifactType {
	if o == nil {
		var ret ArtifactType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SearchedArtifact) GetTypeOk() (*ArtifactType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SearchedArtifact) SetType(v ArtifactType) {
	o.Type = v
}


// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *SearchedArtifact) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchedArtifact) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *SearchedArtifact) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *SearchedArtifact) SetLabels(v []string) {
	o.Labels = &v
}


// GetState returns the State field value
func (o *SearchedArtifact) GetState() ArtifactState {
	if o == nil {
		var ret ArtifactState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SearchedArtifact) GetStateOk() (*ArtifactState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SearchedArtifact) SetState(v ArtifactState) {
	o.State = v
}


// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *SearchedArtifact) GetModifiedOn() string {
	if o == nil || o.ModifiedOn == nil {
		var ret string
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchedArtifact) GetModifiedOnOk() (*string, bool) {
	if o == nil || o.ModifiedOn == nil {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *SearchedArtifact) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn != nil {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given string and assigns it to the ModifiedOn field.
func (o *SearchedArtifact) SetModifiedOn(v string) {
	o.ModifiedOn = &v
}


// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *SearchedArtifact) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchedArtifact) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *SearchedArtifact) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *SearchedArtifact) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}


// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SearchedArtifact) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchedArtifact) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SearchedArtifact) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *SearchedArtifact) SetGroupId(v string) {
	o.GroupId = &v
}


func (o SearchedArtifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["id"] = o.Id
	}
    
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
    
	if true {
		toSerialize["createdOn"] = o.CreatedOn
	}
    
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
    
	if true {
		toSerialize["type"] = o.Type
	}
    
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
    
	if true {
		toSerialize["state"] = o.State
	}
    
	if o.ModifiedOn != nil {
		toSerialize["modifiedOn"] = o.ModifiedOn
	}
    
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
    
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
    
	return json.Marshal(toSerialize)
}

type NullableSearchedArtifact struct {
	value *SearchedArtifact
	isSet bool
}

func (v NullableSearchedArtifact) Get() *SearchedArtifact {
	return v.value
}

func (v *NullableSearchedArtifact) Set(val *SearchedArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchedArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchedArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchedArtifact(val *SearchedArtifact) *NullableSearchedArtifact {
	return &NullableSearchedArtifact{value: val, isSet: true}
}

func (v NullableSearchedArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchedArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

