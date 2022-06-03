# KafkaUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **NullableString** |  | [optional] 
**ReauthenticationEnabled** | Pointer to **NullableBool** | Whether connection reauthentication is enabled or not. If set to true, connection reauthentication on the Kafka instance will be required every 5 minutes. | [optional] 

## Methods

### NewKafkaUpdateRequest

`func NewKafkaUpdateRequest() *KafkaUpdateRequest`

NewKafkaUpdateRequest instantiates a new KafkaUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaUpdateRequestWithDefaults

`func NewKafkaUpdateRequestWithDefaults() *KafkaUpdateRequest`

NewKafkaUpdateRequestWithDefaults instantiates a new KafkaUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *KafkaUpdateRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *KafkaUpdateRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *KafkaUpdateRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *KafkaUpdateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *KafkaUpdateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *KafkaUpdateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetReauthenticationEnabled

`func (o *KafkaUpdateRequest) GetReauthenticationEnabled() bool`

GetReauthenticationEnabled returns the ReauthenticationEnabled field if non-nil, zero value otherwise.

### GetReauthenticationEnabledOk

`func (o *KafkaUpdateRequest) GetReauthenticationEnabledOk() (*bool, bool)`

GetReauthenticationEnabledOk returns a tuple with the ReauthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthenticationEnabled

`func (o *KafkaUpdateRequest) SetReauthenticationEnabled(v bool)`

SetReauthenticationEnabled sets ReauthenticationEnabled field to given value.

### HasReauthenticationEnabled

`func (o *KafkaUpdateRequest) HasReauthenticationEnabled() bool`

HasReauthenticationEnabled returns a boolean if a field has been set.

### SetReauthenticationEnabledNil

`func (o *KafkaUpdateRequest) SetReauthenticationEnabledNil(b bool)`

 SetReauthenticationEnabledNil sets the value for ReauthenticationEnabled to be an explicit nil

### UnsetReauthenticationEnabled
`func (o *KafkaUpdateRequest) UnsetReauthenticationEnabled()`

UnsetReauthenticationEnabled ensures that no value is present for ReauthenticationEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


