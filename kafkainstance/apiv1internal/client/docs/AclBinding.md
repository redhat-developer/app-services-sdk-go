# AclBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the object. Not supported for all object kinds. | [optional] 
**Kind** | Pointer to **string** |  | [optional] [readonly] 
**Href** | Pointer to **string** | Link path to request the object. Not supported for all object kinds. | [optional] 
**ResourceType** | [**AclResourceType**](AclResourceType.md) |  | 
**ResourceName** | **string** |  | 
**PatternType** | [**AclPatternType**](AclPatternType.md) |  | 
**Principal** | **string** | Identifies the user or service account to which an ACL entry is bound. The literal prefix value of &#x60;User:&#x60; is required. May be used to specify all users with value &#x60;User:*&#x60;. | 
**Operation** | [**AclOperation**](AclOperation.md) |  | 
**Permission** | [**AclPermissionType**](AclPermissionType.md) |  | 

## Methods

### NewAclBinding

`func NewAclBinding(resourceType AclResourceType, resourceName string, patternType AclPatternType, principal string, operation AclOperation, permission AclPermissionType, ) *AclBinding`

NewAclBinding instantiates a new AclBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclBindingWithDefaults

`func NewAclBindingWithDefaults() *AclBinding`

NewAclBindingWithDefaults instantiates a new AclBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AclBinding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AclBinding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AclBinding) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AclBinding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AclBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AclBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AclBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AclBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *AclBinding) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AclBinding) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AclBinding) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AclBinding) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetResourceType

`func (o *AclBinding) GetResourceType() AclResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AclBinding) GetResourceTypeOk() (*AclResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AclBinding) SetResourceType(v AclResourceType)`

SetResourceType sets ResourceType field to given value.


### GetResourceName

`func (o *AclBinding) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AclBinding) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AclBinding) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetPatternType

`func (o *AclBinding) GetPatternType() AclPatternType`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *AclBinding) GetPatternTypeOk() (*AclPatternType, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *AclBinding) SetPatternType(v AclPatternType)`

SetPatternType sets PatternType field to given value.


### GetPrincipal

`func (o *AclBinding) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *AclBinding) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *AclBinding) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetOperation

`func (o *AclBinding) GetOperation() AclOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AclBinding) GetOperationOk() (*AclOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AclBinding) SetOperation(v AclOperation)`

SetOperation sets Operation field to given value.


### GetPermission

`func (o *AclBinding) GetPermission() AclPermissionType`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *AclBinding) GetPermissionOk() (*AclPermissionType, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *AclBinding) SetPermission(v AclPermissionType)`

SetPermission sets Permission field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


