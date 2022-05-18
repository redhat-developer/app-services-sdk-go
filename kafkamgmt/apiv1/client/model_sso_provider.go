/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// SsoProvider SSO Provider
type SsoProvider struct {

	Id *string `json:"id,omitempty"`

	Kind *string `json:"kind,omitempty"`

	Href *string `json:"href,omitempty"`

	// name of the sso provider
	Name *string `json:"name,omitempty"`

	// base url
	BaseUrl *string `json:"base_url,omitempty"`

	TokenUrl *string `json:"token_url,omitempty"`

	Jwks *string `json:"jwks,omitempty"`

	ValidIssuer *string `json:"valid_issuer,omitempty"`

}

// NewSsoProvider instantiates a new SsoProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoProvider() *SsoProvider {
	this := SsoProvider{}
	return &this
}

// NewSsoProviderWithDefaults instantiates a new SsoProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoProviderWithDefaults() *SsoProvider {
	this := SsoProvider{}









	return &this
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *SsoProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SsoProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SsoProvider) SetId(v string) {
	o.Id = &v
}


// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SsoProvider) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProvider) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SsoProvider) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *SsoProvider) SetKind(v string) {
	o.Kind = &v
}


// GetHref returns the Href field value if set, zero value otherwise.
func (o *SsoProvider) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProvider) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SsoProvider) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SsoProvider) SetHref(v string) {
	o.Href = &v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *SsoProvider) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProvider) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SsoProvider) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SsoProvider) SetName(v string) {
	o.Name = &v
}


// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *SsoProvider) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProvider) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *SsoProvider) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *SsoProvider) SetBaseUrl(v string) {
	o.BaseUrl = &v
}


// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *SsoProvider) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProvider) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *SsoProvider) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *SsoProvider) SetTokenUrl(v string) {
	o.TokenUrl = &v
}


// GetJwks returns the Jwks field value if set, zero value otherwise.
func (o *SsoProvider) GetJwks() string {
	if o == nil || o.Jwks == nil {
		var ret string
		return ret
	}
	return *o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProvider) GetJwksOk() (*string, bool) {
	if o == nil || o.Jwks == nil {
		return nil, false
	}
	return o.Jwks, true
}

// HasJwks returns a boolean if a field has been set.
func (o *SsoProvider) HasJwks() bool {
	if o != nil && o.Jwks != nil {
		return true
	}

	return false
}

// SetJwks gets a reference to the given string and assigns it to the Jwks field.
func (o *SsoProvider) SetJwks(v string) {
	o.Jwks = &v
}


// GetValidIssuer returns the ValidIssuer field value if set, zero value otherwise.
func (o *SsoProvider) GetValidIssuer() string {
	if o == nil || o.ValidIssuer == nil {
		var ret string
		return ret
	}
	return *o.ValidIssuer
}

// GetValidIssuerOk returns a tuple with the ValidIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoProvider) GetValidIssuerOk() (*string, bool) {
	if o == nil || o.ValidIssuer == nil {
		return nil, false
	}
	return o.ValidIssuer, true
}

// HasValidIssuer returns a boolean if a field has been set.
func (o *SsoProvider) HasValidIssuer() bool {
	if o != nil && o.ValidIssuer != nil {
		return true
	}

	return false
}

// SetValidIssuer gets a reference to the given string and assigns it to the ValidIssuer field.
func (o *SsoProvider) SetValidIssuer(v string) {
	o.ValidIssuer = &v
}


func (o SsoProvider) MarshalJSON() ([]byte, error) {
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
    
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
    
	if o.BaseUrl != nil {
		toSerialize["base_url"] = o.BaseUrl
	}
    
	if o.TokenUrl != nil {
		toSerialize["token_url"] = o.TokenUrl
	}
    
	if o.Jwks != nil {
		toSerialize["jwks"] = o.Jwks
	}
    
	if o.ValidIssuer != nil {
		toSerialize["valid_issuer"] = o.ValidIssuer
	}
    
	return json.Marshal(toSerialize)
}

type NullableSsoProvider struct {
	value *SsoProvider
	isSet bool
}

func (v NullableSsoProvider) Get() *SsoProvider {
	return v.value
}

func (v *NullableSsoProvider) Set(val *SsoProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoProvider(val *SsoProvider) *NullableSsoProvider {
	return &NullableSsoProvider{value: val, isSet: true}
}

func (v NullableSsoProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
