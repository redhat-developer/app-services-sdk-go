# SsoProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** |  | 
**Href** | **string** |  | 
**Name** | Pointer to **string** | name of the sso provider | [optional] 
**BaseUrl** | Pointer to **string** | base url | [optional] 
**TokenUrl** | Pointer to **string** |  | [optional] 
**Jwks** | Pointer to **string** |  | [optional] 
**ValidIssuer** | Pointer to **string** |  | [optional] 

## Methods

### NewSsoProvider

`func NewSsoProvider(id string, kind string, href string, ) *SsoProvider`

NewSsoProvider instantiates a new SsoProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoProviderWithDefaults

`func NewSsoProviderWithDefaults() *SsoProvider`

NewSsoProviderWithDefaults instantiates a new SsoProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SsoProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SsoProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SsoProvider) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *SsoProvider) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SsoProvider) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SsoProvider) SetKind(v string)`

SetKind sets Kind field to given value.


### GetHref

`func (o *SsoProvider) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SsoProvider) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SsoProvider) SetHref(v string)`

SetHref sets Href field to given value.


### GetName

`func (o *SsoProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SsoProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SsoProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SsoProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBaseUrl

`func (o *SsoProvider) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *SsoProvider) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *SsoProvider) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *SsoProvider) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetTokenUrl

`func (o *SsoProvider) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *SsoProvider) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *SsoProvider) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *SsoProvider) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### GetJwks

`func (o *SsoProvider) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *SsoProvider) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *SsoProvider) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *SsoProvider) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetValidIssuer

`func (o *SsoProvider) GetValidIssuer() string`

GetValidIssuer returns the ValidIssuer field if non-nil, zero value otherwise.

### GetValidIssuerOk

`func (o *SsoProvider) GetValidIssuerOk() (*string, bool)`

GetValidIssuerOk returns a tuple with the ValidIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidIssuer

`func (o *SsoProvider) SetValidIssuer(v string)`

SetValidIssuer sets ValidIssuer field to given value.

### HasValidIssuer

`func (o *SsoProvider) HasValidIssuer() bool`

HasValidIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


