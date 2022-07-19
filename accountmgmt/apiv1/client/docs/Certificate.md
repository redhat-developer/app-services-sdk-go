# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cert** | **string** |  | 
**Id** | **string** |  | 
**Key** | **string** |  | 
**Metadata** | **map[string]string** |  | 
**OrganizationId** | **string** |  | 
**Serial** | [**CertificateSerial**](CertificateSerial.md) |  | 

## Methods

### NewCertificate

`func NewCertificate(cert string, id string, key string, metadata map[string]string, organizationId string, serial CertificateSerial, ) *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCert

`func (o *Certificate) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *Certificate) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *Certificate) SetCert(v string)`

SetCert sets Cert field to given value.


### GetId

`func (o *Certificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Certificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Certificate) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *Certificate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Certificate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Certificate) SetKey(v string)`

SetKey sets Key field to given value.


### GetMetadata

`func (o *Certificate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Certificate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Certificate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### GetOrganizationId

`func (o *Certificate) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Certificate) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Certificate) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetSerial

`func (o *Certificate) GetSerial() CertificateSerial`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Certificate) GetSerialOk() (*CertificateSerial, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Certificate) SetSerial(v CertificateSerial)`

SetSerial sets Serial field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


