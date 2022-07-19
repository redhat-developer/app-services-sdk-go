# CertificatesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewCertificatesRequest

`func NewCertificatesRequest(type_ string, ) *CertificatesRequest`

NewCertificatesRequest instantiates a new CertificatesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesRequestWithDefaults

`func NewCertificatesRequestWithDefaults() *CertificatesRequest`

NewCertificatesRequestWithDefaults instantiates a new CertificatesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *CertificatesRequest) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *CertificatesRequest) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *CertificatesRequest) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *CertificatesRequest) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetType

`func (o *CertificatesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificatesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificatesRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


