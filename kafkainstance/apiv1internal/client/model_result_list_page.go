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
)

// ResultListPage struct for ResultListPage
type ResultListPage struct {

	// Total number of entries in the full result set
	Total *float32 `json:"total,omitempty"`

	// Current page number
	Page *int32 `json:"page,omitempty"`

	// Number of entries per page
	Size *float32 `json:"size,omitempty"`

}

// NewResultListPage instantiates a new ResultListPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultListPage() *ResultListPage {
	this := ResultListPage{}
	return &this
}

// NewResultListPageWithDefaults instantiates a new ResultListPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultListPageWithDefaults() *ResultListPage {
	this := ResultListPage{}




	return &this
}


// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ResultListPage) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultListPage) GetTotalOk() (*float32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ResultListPage) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *ResultListPage) SetTotal(v float32) {
	o.Total = &v
}


// GetPage returns the Page field value if set, zero value otherwise.
func (o *ResultListPage) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultListPage) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ResultListPage) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ResultListPage) SetPage(v int32) {
	o.Page = &v
}


// GetSize returns the Size field value if set, zero value otherwise.
func (o *ResultListPage) GetSize() float32 {
	if o == nil || o.Size == nil {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultListPage) GetSizeOk() (*float32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ResultListPage) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *ResultListPage) SetSize(v float32) {
	o.Size = &v
}


func (o ResultListPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
    
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
    
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
    
	return json.Marshal(toSerialize)
}

type NullableResultListPage struct {
	value *ResultListPage
	isSet bool
}

func (v NullableResultListPage) Get() *ResultListPage {
	return v.value
}

func (v *NullableResultListPage) Set(val *ResultListPage) {
	v.value = val
	v.isSet = true
}

func (v NullableResultListPage) IsSet() bool {
	return v.isSet
}

func (v *NullableResultListPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultListPage(val *ResultListPage) *NullableResultListPage {
	return &NullableResultListPage{value: val, isSet: true}
}

func (v NullableResultListPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultListPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

