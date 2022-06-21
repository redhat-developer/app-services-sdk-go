/*
 * Red Hat Openshift SmartEvents Fleet Manager
 *
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * API version: 0.0.1
 * Contact: openbridge-dev@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsclient

import (
	"encoding/json"
	"time"
)

// ProcessorResponse struct for ProcessorResponse
type ProcessorResponse struct {
	Kind *string `json:"kind,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Href *string `json:"href,omitempty"`
	SubmittedAt *time.Time `json:"submitted_at,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	Status *ManagedResourceStatus `json:"status,omitempty"`
	Owner *string `json:"owner,omitempty"`
	Type *ProcessorType `json:"type,omitempty"`
	Filters *[]BaseFilter `json:"filters,omitempty"`
	TransformationTemplate *string `json:"transformationTemplate,omitempty"`
	Action *Action `json:"action,omitempty"`
	Source *Source `json:"source,omitempty"`
}

// NewProcessorResponse instantiates a new ProcessorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorResponse() *ProcessorResponse {
	this := ProcessorResponse{}
	return &this
}

// NewProcessorResponseWithDefaults instantiates a new ProcessorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorResponseWithDefaults() *ProcessorResponse {
	this := ProcessorResponse{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ProcessorResponse) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ProcessorResponse) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ProcessorResponse) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProcessorResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProcessorResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProcessorResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProcessorResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProcessorResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProcessorResponse) SetName(v string) {
	o.Name = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ProcessorResponse) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ProcessorResponse) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ProcessorResponse) SetHref(v string) {
	o.Href = &v
}

// GetSubmittedAt returns the SubmittedAt field value if set, zero value otherwise.
func (o *ProcessorResponse) GetSubmittedAt() time.Time {
	if o == nil || o.SubmittedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetSubmittedAtOk() (*time.Time, bool) {
	if o == nil || o.SubmittedAt == nil {
		return nil, false
	}
	return o.SubmittedAt, true
}

// HasSubmittedAt returns a boolean if a field has been set.
func (o *ProcessorResponse) HasSubmittedAt() bool {
	if o != nil && o.SubmittedAt != nil {
		return true
	}

	return false
}

// SetSubmittedAt gets a reference to the given time.Time and assigns it to the SubmittedAt field.
func (o *ProcessorResponse) SetSubmittedAt(v time.Time) {
	o.SubmittedAt = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *ProcessorResponse) GetPublishedAt() time.Time {
	if o == nil || o.PublishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || o.PublishedAt == nil {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *ProcessorResponse) HasPublishedAt() bool {
	if o != nil && o.PublishedAt != nil {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *ProcessorResponse) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProcessorResponse) GetStatus() ManagedResourceStatus {
	if o == nil || o.Status == nil {
		var ret ManagedResourceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetStatusOk() (*ManagedResourceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProcessorResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ManagedResourceStatus and assigns it to the Status field.
func (o *ProcessorResponse) SetStatus(v ManagedResourceStatus) {
	o.Status = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ProcessorResponse) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ProcessorResponse) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ProcessorResponse) SetOwner(v string) {
	o.Owner = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProcessorResponse) GetType() ProcessorType {
	if o == nil || o.Type == nil {
		var ret ProcessorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetTypeOk() (*ProcessorType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProcessorResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ProcessorType and assigns it to the Type field.
func (o *ProcessorResponse) SetType(v ProcessorType) {
	o.Type = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ProcessorResponse) GetFilters() []BaseFilter {
	if o == nil || o.Filters == nil {
		var ret []BaseFilter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetFiltersOk() (*[]BaseFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ProcessorResponse) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []BaseFilter and assigns it to the Filters field.
func (o *ProcessorResponse) SetFilters(v []BaseFilter) {
	o.Filters = &v
}

// GetTransformationTemplate returns the TransformationTemplate field value if set, zero value otherwise.
func (o *ProcessorResponse) GetTransformationTemplate() string {
	if o == nil || o.TransformationTemplate == nil {
		var ret string
		return ret
	}
	return *o.TransformationTemplate
}

// GetTransformationTemplateOk returns a tuple with the TransformationTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetTransformationTemplateOk() (*string, bool) {
	if o == nil || o.TransformationTemplate == nil {
		return nil, false
	}
	return o.TransformationTemplate, true
}

// HasTransformationTemplate returns a boolean if a field has been set.
func (o *ProcessorResponse) HasTransformationTemplate() bool {
	if o != nil && o.TransformationTemplate != nil {
		return true
	}

	return false
}

// SetTransformationTemplate gets a reference to the given string and assigns it to the TransformationTemplate field.
func (o *ProcessorResponse) SetTransformationTemplate(v string) {
	o.TransformationTemplate = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ProcessorResponse) GetAction() Action {
	if o == nil || o.Action == nil {
		var ret Action
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetActionOk() (*Action, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ProcessorResponse) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given Action and assigns it to the Action field.
func (o *ProcessorResponse) SetAction(v Action) {
	o.Action = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProcessorResponse) GetSource() Source {
	if o == nil || o.Source == nil {
		var ret Source
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorResponse) GetSourceOk() (*Source, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProcessorResponse) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given Source and assigns it to the Source field.
func (o *ProcessorResponse) SetSource(v Source) {
	o.Source = &v
}

func (o ProcessorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.SubmittedAt != nil {
		toSerialize["submitted_at"] = o.SubmittedAt
	}
	if o.PublishedAt != nil {
		toSerialize["published_at"] = o.PublishedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.TransformationTemplate != nil {
		toSerialize["transformationTemplate"] = o.TransformationTemplate
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorResponse struct {
	value *ProcessorResponse
	isSet bool
}

func (v NullableProcessorResponse) Get() *ProcessorResponse {
	return v.value
}

func (v *NullableProcessorResponse) Set(val *ProcessorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorResponse(val *ProcessorResponse) *NullableProcessorResponse {
	return &NullableProcessorResponse{value: val, isSet: true}
}

func (v NullableProcessorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


