# SupportedKafkaInstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Indicates the type of this object. Will be &#39;SupportedKafkaInstanceType&#39; link. | [optional] 
**Id** | Pointer to **string** | Unique identifier of the Kafka instance type. | [optional] 
**Sizes** | Pointer to [**[]SupportedKafkaSize**](SupportedKafkaSize.md) |  A list of Kafka instance sizes available for this instance type | [optional] 


## Methods

### NewSupportedKafkaInstanceType

`func NewSupportedKafkaInstanceType() *SupportedKafkaInstanceType`

NewSupportedKafkaInstanceType instantiates a new SupportedKafkaInstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedKafkaInstanceTypeWithDefaults

`func NewSupportedKafkaInstanceTypeWithDefaults() *SupportedKafkaInstanceType`

NewSupportedKafkaInstanceTypeWithDefaults instantiates a new SupportedKafkaInstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetKind

`func (o *SupportedKafkaInstanceType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SupportedKafkaInstanceType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SupportedKafkaInstanceType) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SupportedKafkaInstanceType) HasKind() bool`

HasKind returns a boolean if a field has been set.


### GetId

`func (o *SupportedKafkaInstanceType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportedKafkaInstanceType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportedKafkaInstanceType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportedKafkaInstanceType) HasId() bool`

HasId returns a boolean if a field has been set.


### GetSizes

`func (o *SupportedKafkaInstanceType) GetSizes() []SupportedKafkaSize`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *SupportedKafkaInstanceType) GetSizesOk() (*[]SupportedKafkaSize, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *SupportedKafkaInstanceType) SetSizes(v []SupportedKafkaSize)`

SetSizes sets Sizes field to given value.

### HasSizes

`func (o *SupportedKafkaInstanceType) HasSizes() bool`

HasSizes returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

