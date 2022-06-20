# AclBindingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | [**AclResourceType**](AclResourceType.md) |  | 
**ResourceName** | **string** |  | 
**PatternType** | [**AclPatternType**](AclPatternType.md) |  | 
**Principal** | **string** | Identifies the user or service account to which an ACL entry is bound. The literal prefix value of &#x60;User:&#x60; is required. May be used to specify all users with value &#x60;User:*&#x60;. | 
**Operation** | [**AclOperation**](AclOperation.md) |  | 
**Permission** | [**AclPermissionType**](AclPermissionType.md) |  | 

## Methods

### NewAclBindingAllOf

`func NewAclBindingAllOf(resourceType AclResourceType, resourceName string, patternType AclPatternType, principal string, operation AclOperation, permission AclPermissionType, ) *AclBindingAllOf`

NewAclBindingAllOf instantiates a new AclBindingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclBindingAllOfWithDefaults

`func NewAclBindingAllOfWithDefaults() *AclBindingAllOf`

NewAclBindingAllOfWithDefaults instantiates a new AclBindingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *AclBindingAllOf) GetResourceType() AclResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AclBindingAllOf) GetResourceTypeOk() (*AclResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AclBindingAllOf) SetResourceType(v AclResourceType)`

SetResourceType sets ResourceType field to given value.


### GetResourceName

`func (o *AclBindingAllOf) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AclBindingAllOf) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AclBindingAllOf) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetPatternType

`func (o *AclBindingAllOf) GetPatternType() AclPatternType`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *AclBindingAllOf) GetPatternTypeOk() (*AclPatternType, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *AclBindingAllOf) SetPatternType(v AclPatternType)`

SetPatternType sets PatternType field to given value.


### GetPrincipal

`func (o *AclBindingAllOf) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *AclBindingAllOf) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *AclBindingAllOf) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetOperation

`func (o *AclBindingAllOf) GetOperation() AclOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AclBindingAllOf) GetOperationOk() (*AclOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AclBindingAllOf) SetOperation(v AclOperation)`

SetOperation sets Operation field to given value.


### GetPermission

`func (o *AclBindingAllOf) GetPermission() AclPermissionType`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *AclBindingAllOf) GetPermissionOk() (*AclPermissionType, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *AclBindingAllOf) SetPermission(v AclPermissionType)`

SetPermission sets Permission field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


