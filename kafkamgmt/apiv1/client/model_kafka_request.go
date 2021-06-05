/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kakfa instances and connectors.
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
	"time"
)

// KafkaRequest struct for KafkaRequest
type KafkaRequest struct {

	Id *string `json:"id,omitempty"`

	Kind *string `json:"kind,omitempty"`

	Href *string `json:"href,omitempty"`

	Status *string `json:"status,omitempty"`

	CloudProvider *string `json:"cloud_provider,omitempty"`

	MultiAz *bool `json:"multi_az,omitempty"`

	Region *string `json:"region,omitempty"`

	Owner *string `json:"owner,omitempty"`

	Name *string `json:"name,omitempty"`

	BootstrapServerHost *string `json:"bootstrap_server_host,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`

	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	FailedReason *string `json:"failed_reason,omitempty"`

	Version *string `json:"version,omitempty"`

}

// NewKafkaRequest instantiates a new KafkaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaRequest() *KafkaRequest {
	this := KafkaRequest{}
	return &this
}

// NewKafkaRequestWithDefaults instantiates a new KafkaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaRequestWithDefaults() *KafkaRequest {
	this := KafkaRequest{}















	return &this
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *KafkaRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KafkaRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KafkaRequest) SetId(v string) {
	o.Id = &v
}


// GetKind returns the Kind field value if set, zero value otherwise.
func (o *KafkaRequest) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *KafkaRequest) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *KafkaRequest) SetKind(v string) {
	o.Kind = &v
}


// GetHref returns the Href field value if set, zero value otherwise.
func (o *KafkaRequest) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *KafkaRequest) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *KafkaRequest) SetHref(v string) {
	o.Href = &v
}


// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KafkaRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KafkaRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *KafkaRequest) SetStatus(v string) {
	o.Status = &v
}


// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *KafkaRequest) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *KafkaRequest) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *KafkaRequest) SetCloudProvider(v string) {
	o.CloudProvider = &v
}


// GetMultiAz returns the MultiAz field value if set, zero value otherwise.
func (o *KafkaRequest) GetMultiAz() bool {
	if o == nil || o.MultiAz == nil {
		var ret bool
		return ret
	}
	return *o.MultiAz
}

// GetMultiAzOk returns a tuple with the MultiAz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetMultiAzOk() (*bool, bool) {
	if o == nil || o.MultiAz == nil {
		return nil, false
	}
	return o.MultiAz, true
}

// HasMultiAz returns a boolean if a field has been set.
func (o *KafkaRequest) HasMultiAz() bool {
	if o != nil && o.MultiAz != nil {
		return true
	}

	return false
}

// SetMultiAz gets a reference to the given bool and assigns it to the MultiAz field.
func (o *KafkaRequest) SetMultiAz(v bool) {
	o.MultiAz = &v
}


// GetRegion returns the Region field value if set, zero value otherwise.
func (o *KafkaRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *KafkaRequest) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *KafkaRequest) SetRegion(v string) {
	o.Region = &v
}


// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *KafkaRequest) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *KafkaRequest) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *KafkaRequest) SetOwner(v string) {
	o.Owner = &v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *KafkaRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KafkaRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KafkaRequest) SetName(v string) {
	o.Name = &v
}


// GetBootstrapServerHost returns the BootstrapServerHost field value if set, zero value otherwise.
func (o *KafkaRequest) GetBootstrapServerHost() string {
	if o == nil || o.BootstrapServerHost == nil {
		var ret string
		return ret
	}
	return *o.BootstrapServerHost
}

// GetBootstrapServerHostOk returns a tuple with the BootstrapServerHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetBootstrapServerHostOk() (*string, bool) {
	if o == nil || o.BootstrapServerHost == nil {
		return nil, false
	}
	return o.BootstrapServerHost, true
}

// HasBootstrapServerHost returns a boolean if a field has been set.
func (o *KafkaRequest) HasBootstrapServerHost() bool {
	if o != nil && o.BootstrapServerHost != nil {
		return true
	}

	return false
}

// SetBootstrapServerHost gets a reference to the given string and assigns it to the BootstrapServerHost field.
func (o *KafkaRequest) SetBootstrapServerHost(v string) {
	o.BootstrapServerHost = &v
}


// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KafkaRequest) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KafkaRequest) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *KafkaRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}


// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KafkaRequest) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KafkaRequest) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *KafkaRequest) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}


// GetFailedReason returns the FailedReason field value if set, zero value otherwise.
func (o *KafkaRequest) GetFailedReason() string {
	if o == nil || o.FailedReason == nil {
		var ret string
		return ret
	}
	return *o.FailedReason
}

// GetFailedReasonOk returns a tuple with the FailedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetFailedReasonOk() (*string, bool) {
	if o == nil || o.FailedReason == nil {
		return nil, false
	}
	return o.FailedReason, true
}

// HasFailedReason returns a boolean if a field has been set.
func (o *KafkaRequest) HasFailedReason() bool {
	if o != nil && o.FailedReason != nil {
		return true
	}

	return false
}

// SetFailedReason gets a reference to the given string and assigns it to the FailedReason field.
func (o *KafkaRequest) SetFailedReason(v string) {
	o.FailedReason = &v
}


// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KafkaRequest) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaRequest) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KafkaRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *KafkaRequest) SetVersion(v string) {
	o.Version = &v
}


func (o KafkaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
    
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
    
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
    
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
    
	if o.CloudProvider != nil {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
    
	if o.MultiAz != nil {
		toSerialize["multi_az"] = o.MultiAz
	}
    
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
    
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
    
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.BootstrapServerHost != nil {
		toSerialize["bootstrap_server_host"] = o.BootstrapServerHost
	}
    
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
    
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
    
	if o.FailedReason != nil {
		toSerialize["failed_reason"] = o.FailedReason
	}
    
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
    
	return json.Marshal(toSerialize)
}

type NullableKafkaRequest struct {
	value *KafkaRequest
	isSet bool
}

func (v NullableKafkaRequest) Get() *KafkaRequest {
	return v.value
}

func (v *NullableKafkaRequest) Set(val *KafkaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaRequest(val *KafkaRequest) *NullableKafkaRequest {
	return &NullableKafkaRequest{value: val, isSet: true}
}

func (v NullableKafkaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

