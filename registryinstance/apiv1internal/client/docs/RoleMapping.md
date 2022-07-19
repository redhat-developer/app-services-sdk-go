# RoleMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalId** | **string** |  | 
**Role** | [**RoleType**](RoleType.md) |  | 
**PrincipalName** | Pointer to **string** | A friendly name for the principal. | [optional] 

## Methods

### NewRoleMapping

`func NewRoleMapping(principalId string, role RoleType, ) *RoleMapping`

NewRoleMapping instantiates a new RoleMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMappingWithDefaults

`func NewRoleMappingWithDefaults() *RoleMapping`

NewRoleMappingWithDefaults instantiates a new RoleMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalId

`func (o *RoleMapping) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *RoleMapping) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *RoleMapping) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetRole

`func (o *RoleMapping) GetRole() RoleType`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleMapping) GetRoleOk() (*RoleType, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleMapping) SetRole(v RoleType)`

SetRole sets Role field to given value.


### GetPrincipalName

`func (o *RoleMapping) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *RoleMapping) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *RoleMapping) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *RoleMapping) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


