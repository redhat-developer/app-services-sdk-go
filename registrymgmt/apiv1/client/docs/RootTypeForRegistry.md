# RootTypeForRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | [**RegistryStatusValue**](RegistryStatusValue.md) |  | 
**RegistryUrl** | Pointer to **string** |  | [optional] 
**BrowserUrl** | Pointer to **string** |  | [optional] 
**Name** | **string** | User-defined Registry instance name. Does not have to be unique. | 
**RegistryDeploymentId** | Pointer to **int32** | Identifier of a multi-tenant deployment, where this Service Registry instance resides. | [optional] 
**Owner** | **string** | Registry instance owner. | 
**Description** | Pointer to **string** | Description of the Registry instance. | [optional] 
**CreatedAt** | **time.Time** | ISO 8601 UTC timestamp. | 
**UpdatedAt** | **time.Time** | ISO 8601 UTC timestamp. | 
**InstanceType** | [**RegistryInstanceTypeValue**](RegistryInstanceTypeValue.md) |  | 

## Methods

### NewRootTypeForRegistry

`func NewRootTypeForRegistry(id string, status RegistryStatusValue, name string, owner string, createdAt time.Time, updatedAt time.Time, instanceType RegistryInstanceTypeValue, ) *RootTypeForRegistry`

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

`func (o *RootTypeForRegistry) GetStatus() RegistryStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RootTypeForRegistry) GetStatusOk() (*RegistryStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RootTypeForRegistry) SetStatus(v RegistryStatusValue)`

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

### HasRegistryUrl

`func (o *RootTypeForRegistry) HasRegistryUrl() bool`

HasRegistryUrl returns a boolean if a field has been set.

### GetBrowserUrl

`func (o *RootTypeForRegistry) GetBrowserUrl() string`

GetBrowserUrl returns the BrowserUrl field if non-nil, zero value otherwise.

### GetBrowserUrlOk

`func (o *RootTypeForRegistry) GetBrowserUrlOk() (*string, bool)`

GetBrowserUrlOk returns a tuple with the BrowserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserUrl

`func (o *RootTypeForRegistry) SetBrowserUrl(v string)`

SetBrowserUrl sets BrowserUrl field to given value.

### HasBrowserUrl

`func (o *RootTypeForRegistry) HasBrowserUrl() bool`

HasBrowserUrl returns a boolean if a field has been set.

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

### GetOwner

`func (o *RootTypeForRegistry) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RootTypeForRegistry) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RootTypeForRegistry) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetDescription

`func (o *RootTypeForRegistry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RootTypeForRegistry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RootTypeForRegistry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RootTypeForRegistry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RootTypeForRegistry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RootTypeForRegistry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RootTypeForRegistry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RootTypeForRegistry) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RootTypeForRegistry) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RootTypeForRegistry) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetInstanceType

`func (o *RootTypeForRegistry) GetInstanceType() RegistryInstanceTypeValue`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *RootTypeForRegistry) GetInstanceTypeOk() (*RegistryInstanceTypeValue, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *RootTypeForRegistry) SetInstanceType(v RegistryInstanceTypeValue)`

SetInstanceType sets InstanceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


