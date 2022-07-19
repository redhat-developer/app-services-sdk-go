# CertificateSerial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**Expiration** | **time.Time** |  | 
**Id** | **int64** |  | 
**Serial** | **int64** |  | 
**Updated** | **time.Time** |  | 

## Methods

### NewCertificateSerial

`func NewCertificateSerial(created time.Time, expiration time.Time, id int64, serial int64, updated time.Time, ) *CertificateSerial`

NewCertificateSerial instantiates a new CertificateSerial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateSerialWithDefaults

`func NewCertificateSerialWithDefaults() *CertificateSerial`

NewCertificateSerialWithDefaults instantiates a new CertificateSerial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CertificateSerial) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CertificateSerial) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CertificateSerial) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExpiration

`func (o *CertificateSerial) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CertificateSerial) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CertificateSerial) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.


### GetId

`func (o *CertificateSerial) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateSerial) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateSerial) SetId(v int64)`

SetId sets Id field to given value.


### GetSerial

`func (o *CertificateSerial) GetSerial() int64`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CertificateSerial) GetSerialOk() (*int64, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CertificateSerial) SetSerial(v int64)`

SetSerial sets Serial field to given value.


### GetUpdated

`func (o *CertificateSerial) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CertificateSerial) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CertificateSerial) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


