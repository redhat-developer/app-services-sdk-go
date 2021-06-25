# RootTypeForRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | [**RegistryStatusValueRest**](RegistryStatusValueRest.md) |  | 
**RegistryUrl** | **string** |  | 
**Name** | Pointer to **string** | User-defined Registry name. Does not have to be unique. | [optional] 
**RegistryDeploymentId** | Pointer to **int32** | Identifier of a multi-tenant deployment, where this Service Registry instance resides. | [optional] 


## Methods

### NewRootTypeForRegistry

`func NewRootTypeForRegistry(id string, status RegistryStatusValueRest, registryUrl string, ) *RootTypeForRegistry`

NewRootTypeForRegistry instantiates a new RootTypeForRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootTypeForRegistryWithDefaults

`func NewRootTypeForRegistryWithDefaults() *RootTypeForRegistry`

NewRootTypeForRegistryWithDefaults instantiates a new RootTypeForRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetId

`func (o *RootTypeForRegistry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RootTypeForRegistry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RootTypeForRegistry) SetId(v string)`

SetId sets Id field to given value.



### GetStatus

`func (o *RootTypeForRegistry) GetStatus() RegistryStatusValueRest`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RootTypeForRegistry) GetStatusOk() (*RegistryStatusValueRest, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RootTypeForRegistry) SetStatus(v RegistryStatusValueRest)`

SetStatus sets Status field to given value.



### GetRegistryUrl

`func (o *RootTypeForRegistry) GetRegistryUrl() string`

GetRegistryUrl returns the RegistryUrl field if non-nil, zero value otherwise.

### GetRegistryUrlOk

`func (o *RootTypeForRegistry) GetRegistryUrlOk() (*string, bool)`

GetRegistryUrlOk returns a tuple with the RegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryUrl

`func (o *RootTypeForRegistry) SetRegistryUrl(v string)`

SetRegistryUrl sets RegistryUrl field to given value.



### GetName

`func (o *RootTypeForRegistry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RootTypeForRegistry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RootTypeForRegistry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RootTypeForRegistry) HasName() bool`

HasName returns a boolean if a field has been set.


### GetRegistryDeploymentId

`func (o *RootTypeForRegistry) GetRegistryDeploymentId() int32`

GetRegistryDeploymentId returns the RegistryDeploymentId field if non-nil, zero value otherwise.

### GetRegistryDeploymentIdOk

`func (o *RootTypeForRegistry) GetRegistryDeploymentIdOk() (*int32, bool)`

GetRegistryDeploymentIdOk returns a tuple with the RegistryDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryDeploymentId

`func (o *RootTypeForRegistry) SetRegistryDeploymentId(v int32)`

SetRegistryDeploymentId sets RegistryDeploymentId field to given value.

### HasRegistryDeploymentId

`func (o *RootTypeForRegistry) HasRegistryDeploymentId() bool`

HasRegistryDeploymentId returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

