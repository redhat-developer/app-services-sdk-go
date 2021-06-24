/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// ConsumerGroupList A list of consumer groups
type ConsumerGroupList struct {

	// Consumer group list items
	Items *[]ConsumerGroup `json:"items,omitempty"`

	// The total number of consumer groups.
	Total *float32 `json:"total,omitempty"`

	// The number of consumer groups per page.
	Size *float32 `json:"size,omitempty"`

	// The page
	Page *int32 `json:"page,omitempty"`

}

// NewConsumerGroupList instantiates a new ConsumerGroupList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerGroupList() *ConsumerGroupList {
	this := ConsumerGroupList{}
	return &this
}

// NewConsumerGroupListWithDefaults instantiates a new ConsumerGroupList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerGroupListWithDefaults() *ConsumerGroupList {
	this := ConsumerGroupList{}





	return &this
}


// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConsumerGroupList) GetItems() []ConsumerGroup {
	if o == nil || o.Items == nil {
		var ret []ConsumerGroup
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupList) GetItemsOk() (*[]ConsumerGroup, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConsumerGroupList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ConsumerGroup and assigns it to the Items field.
func (o *ConsumerGroupList) SetItems(v []ConsumerGroup) {
	o.Items = &v
}


// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ConsumerGroupList) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupList) GetTotalOk() (*float32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ConsumerGroupList) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *ConsumerGroupList) SetTotal(v float32) {
	o.Total = &v
}


// GetSize returns the Size field value if set, zero value otherwise.
func (o *ConsumerGroupList) GetSize() float32 {
	if o == nil || o.Size == nil {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupList) GetSizeOk() (*float32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ConsumerGroupList) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *ConsumerGroupList) SetSize(v float32) {
	o.Size = &v
}


// GetPage returns the Page field value if set, zero value otherwise.
func (o *ConsumerGroupList) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupList) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ConsumerGroupList) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ConsumerGroupList) SetPage(v int32) {
	o.Page = &v
}


func (o ConsumerGroupList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
    
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
    
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
    
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
    
	return json.Marshal(toSerialize)
}

type NullableConsumerGroupList struct {
	value *ConsumerGroupList
	isSet bool
}

func (v NullableConsumerGroupList) Get() *ConsumerGroupList {
	return v.value
}

func (v *NullableConsumerGroupList) Set(val *ConsumerGroupList) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroupList) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroupList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroupList(val *ConsumerGroupList) *NullableConsumerGroupList {
	return &NullableConsumerGroupList{value: val, isSet: true}
}

func (v NullableConsumerGroupList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerGroupList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

