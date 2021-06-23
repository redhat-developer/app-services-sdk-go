# KafkaConnectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapServer** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 


## Methods

### NewKafkaConnectionSettings

`func NewKafkaConnectionSettings() *KafkaConnectionSettings`

NewKafkaConnectionSettings instantiates a new KafkaConnectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaConnectionSettingsWithDefaults

`func NewKafkaConnectionSettingsWithDefaults() *KafkaConnectionSettings`

NewKafkaConnectionSettingsWithDefaults instantiates a new KafkaConnectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetBootstrapServer

`func (o *KafkaConnectionSettings) GetBootstrapServer() string`

GetBootstrapServer returns the BootstrapServer field if non-nil, zero value otherwise.

### GetBootstrapServerOk

`func (o *KafkaConnectionSettings) GetBootstrapServerOk() (*string, bool)`

GetBootstrapServerOk returns a tuple with the BootstrapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServer

`func (o *KafkaConnectionSettings) SetBootstrapServer(v string)`

SetBootstrapServer sets BootstrapServer field to given value.

### HasBootstrapServer

`func (o *KafkaConnectionSettings) HasBootstrapServer() bool`

HasBootstrapServer returns a boolean if a field has been set.


### GetClientId

`func (o *KafkaConnectionSettings) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *KafkaConnectionSettings) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *KafkaConnectionSettings) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *KafkaConnectionSettings) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


### GetClientSecret

`func (o *KafkaConnectionSettings) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *KafkaConnectionSettings) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *KafkaConnectionSettings) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *KafkaConnectionSettings) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

