# RegistryRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Status** | [**RegistryStatusValueRest**](RegistryStatusValueRest.md) |  | 
**RegistryUrl** | **string** |  | 
**Name** | Pointer to **string** | User-defined Registry name. Does not have to be unique. | [optional] 
**RegistryDeploymentId** | Pointer to **int32** | Identifier of a multi-tenant deployment, where this Service Registry instance resides. | [optional] 
**Owner** | Pointer to **string** | Registry instance owner | [optional] 


## Methods

### NewRegistryRest

`func NewRegistryRest(id string, status RegistryStatusValueRest, registryUrl string, ) *RegistryRest`

NewRegistryRest instantiates a new RegistryRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryRestWithDefaults

`func NewRegistryRestWithDefaults() *RegistryRest`

NewRegistryRestWithDefaults instantiates a new RegistryRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetId

`func (o *RegistryRest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryRest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryRest) SetId(v string)`

SetId sets Id field to given value.



### GetKind

`func (o *RegistryRest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RegistryRest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RegistryRest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RegistryRest) HasKind() bool`

HasKind returns a boolean if a field has been set.


### GetHref

`func (o *RegistryRest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RegistryRest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RegistryRest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RegistryRest) HasHref() bool`

HasHref returns a boolean if a field has been set.


### GetStatus

`func (o *RegistryRest) GetStatus() RegistryStatusValueRest`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegistryRest) GetStatusOk() (*RegistryStatusValueRest, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegistryRest) SetStatus(v RegistryStatusValueRest)`

SetStatus sets Status field to given value.



### GetRegistryUrl

`func (o *RegistryRest) GetRegistryUrl() string`

GetRegistryUrl returns the RegistryUrl field if non-nil, zero value otherwise.

### GetRegistryUrlOk

`func (o *RegistryRest) GetRegistryUrlOk() (*string, bool)`

GetRegistryUrlOk returns a tuple with the RegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryUrl

`func (o *RegistryRest) SetRegistryUrl(v string)`

SetRegistryUrl sets RegistryUrl field to given value.



### GetName

`func (o *RegistryRest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryRest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryRest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegistryRest) HasName() bool`

HasName returns a boolean if a field has been set.


### GetRegistryDeploymentId

`func (o *RegistryRest) GetRegistryDeploymentId() int32`

GetRegistryDeploymentId returns the RegistryDeploymentId field if non-nil, zero value otherwise.

### GetRegistryDeploymentIdOk

`func (o *RegistryRest) GetRegistryDeploymentIdOk() (*int32, bool)`

GetRegistryDeploymentIdOk returns a tuple with the RegistryDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryDeploymentId

`func (o *RegistryRest) SetRegistryDeploymentId(v int32)`

SetRegistryDeploymentId sets RegistryDeploymentId field to given value.

### HasRegistryDeploymentId

`func (o *RegistryRest) HasRegistryDeploymentId() bool`

HasRegistryDeploymentId returns a boolean if a field has been set.


### GetOwner

`func (o *RegistryRest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RegistryRest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RegistryRest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *RegistryRest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

