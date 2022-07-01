# ServiceAccountListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | server generated unique id of the service account | 
**Kind** | **string** |  | 
**Href** | **string** |  | 
**ClientId** | Pointer to **string** | client id of the service account | [optional] 
**Name** | Pointer to **string** | name of the service account | [optional] 
**Owner** | Pointer to **string** | owner of the service account | [optional] 
**CreatedBy** | Pointer to **string** | service account created by the user | [optional] 
**CreatedAt** | Pointer to **time.Time** | service account creation timestamp | [optional] 
**Description** | Pointer to **string** | description of the service account | [optional] 

## Methods

### NewServiceAccountListItem

`func NewServiceAccountListItem(id string, kind string, href string, ) *ServiceAccountListItem`

NewServiceAccountListItem instantiates a new ServiceAccountListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountListItemWithDefaults

`func NewServiceAccountListItemWithDefaults() *ServiceAccountListItem`

NewServiceAccountListItemWithDefaults instantiates a new ServiceAccountListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceAccountListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountListItem) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ServiceAccountListItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceAccountListItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceAccountListItem) SetKind(v string)`

SetKind sets Kind field to given value.


### GetHref

`func (o *ServiceAccountListItem) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ServiceAccountListItem) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ServiceAccountListItem) SetHref(v string)`

SetHref sets Href field to given value.


### GetClientId

`func (o *ServiceAccountListItem) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ServiceAccountListItem) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ServiceAccountListItem) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ServiceAccountListItem) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetName

`func (o *ServiceAccountListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountListItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceAccountListItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *ServiceAccountListItem) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ServiceAccountListItem) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ServiceAccountListItem) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ServiceAccountListItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ServiceAccountListItem) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServiceAccountListItem) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServiceAccountListItem) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ServiceAccountListItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceAccountListItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceAccountListItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceAccountListItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceAccountListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceAccountListItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceAccountListItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceAccountListItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceAccountListItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


