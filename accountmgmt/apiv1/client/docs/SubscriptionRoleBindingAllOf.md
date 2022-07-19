# SubscriptionRoleBindingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AccountReference**](AccountReference.md) |  | [optional] 
**AccountEmail** | Pointer to **string** |  | [optional] 
**AccountUsername** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Role** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**Subscription** | Pointer to [**ObjectReference**](ObjectReference.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSubscriptionRoleBindingAllOf

`func NewSubscriptionRoleBindingAllOf() *SubscriptionRoleBindingAllOf`

NewSubscriptionRoleBindingAllOf instantiates a new SubscriptionRoleBindingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionRoleBindingAllOfWithDefaults

`func NewSubscriptionRoleBindingAllOfWithDefaults() *SubscriptionRoleBindingAllOf`

NewSubscriptionRoleBindingAllOfWithDefaults instantiates a new SubscriptionRoleBindingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SubscriptionRoleBindingAllOf) GetAccount() AccountReference`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SubscriptionRoleBindingAllOf) GetAccountOk() (*AccountReference, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SubscriptionRoleBindingAllOf) SetAccount(v AccountReference)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SubscriptionRoleBindingAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountEmail

`func (o *SubscriptionRoleBindingAllOf) GetAccountEmail() string`

GetAccountEmail returns the AccountEmail field if non-nil, zero value otherwise.

### GetAccountEmailOk

`func (o *SubscriptionRoleBindingAllOf) GetAccountEmailOk() (*string, bool)`

GetAccountEmailOk returns a tuple with the AccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEmail

`func (o *SubscriptionRoleBindingAllOf) SetAccountEmail(v string)`

SetAccountEmail sets AccountEmail field to given value.

### HasAccountEmail

`func (o *SubscriptionRoleBindingAllOf) HasAccountEmail() bool`

HasAccountEmail returns a boolean if a field has been set.

### GetAccountUsername

`func (o *SubscriptionRoleBindingAllOf) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *SubscriptionRoleBindingAllOf) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *SubscriptionRoleBindingAllOf) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *SubscriptionRoleBindingAllOf) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionRoleBindingAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionRoleBindingAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionRoleBindingAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionRoleBindingAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRole

`func (o *SubscriptionRoleBindingAllOf) GetRole() ObjectReference`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SubscriptionRoleBindingAllOf) GetRoleOk() (*ObjectReference, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SubscriptionRoleBindingAllOf) SetRole(v ObjectReference)`

SetRole sets Role field to given value.

### HasRole

`func (o *SubscriptionRoleBindingAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSubscription

`func (o *SubscriptionRoleBindingAllOf) GetSubscription() ObjectReference`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SubscriptionRoleBindingAllOf) GetSubscriptionOk() (*ObjectReference, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SubscriptionRoleBindingAllOf) SetSubscription(v ObjectReference)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *SubscriptionRoleBindingAllOf) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionRoleBindingAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionRoleBindingAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionRoleBindingAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriptionRoleBindingAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


