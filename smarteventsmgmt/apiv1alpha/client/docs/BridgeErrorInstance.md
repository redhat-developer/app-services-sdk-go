# BridgeErrorInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**BridgeErrorType**](BridgeErrorType.md) |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewBridgeErrorInstance

`func NewBridgeErrorInstance() *BridgeErrorInstance`

NewBridgeErrorInstance instantiates a new BridgeErrorInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBridgeErrorInstanceWithDefaults

`func NewBridgeErrorInstanceWithDefaults() *BridgeErrorInstance`

NewBridgeErrorInstanceWithDefaults instantiates a new BridgeErrorInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BridgeErrorInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BridgeErrorInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BridgeErrorInstance) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BridgeErrorInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *BridgeErrorInstance) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BridgeErrorInstance) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BridgeErrorInstance) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BridgeErrorInstance) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *BridgeErrorInstance) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BridgeErrorInstance) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BridgeErrorInstance) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BridgeErrorInstance) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetType

`func (o *BridgeErrorInstance) GetType() BridgeErrorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BridgeErrorInstance) GetTypeOk() (*BridgeErrorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BridgeErrorInstance) SetType(v BridgeErrorType)`

SetType sets Type field to given value.

### HasType

`func (o *BridgeErrorInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *BridgeErrorInstance) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BridgeErrorInstance) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BridgeErrorInstance) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *BridgeErrorInstance) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


