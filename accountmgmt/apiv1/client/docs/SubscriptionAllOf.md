# SubscriptionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**[]Capability**](Capability.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Creator** | Pointer to [**AccountReference**](AccountReference.md) |  | [optional] 
**EvalExpirationDate** | Pointer to **time.Time** | Calulated as the subscription created date + 60 days | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) |  | [optional] 
**Metrics** | Pointer to [**[]OneMetric**](OneMetric.md) |  | [optional] 
**NotificationContacts** | Pointer to [**[]Account**](Account.md) |  | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSubscriptionAllOf

`func NewSubscriptionAllOf() *SubscriptionAllOf`

NewSubscriptionAllOf instantiates a new SubscriptionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAllOfWithDefaults

`func NewSubscriptionAllOfWithDefaults() *SubscriptionAllOf`

NewSubscriptionAllOfWithDefaults instantiates a new SubscriptionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *SubscriptionAllOf) GetCapabilities() []Capability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *SubscriptionAllOf) GetCapabilitiesOk() (*[]Capability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *SubscriptionAllOf) SetCapabilities(v []Capability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *SubscriptionAllOf) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionAllOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *SubscriptionAllOf) GetCreator() AccountReference`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SubscriptionAllOf) GetCreatorOk() (*AccountReference, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SubscriptionAllOf) SetCreator(v AccountReference)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *SubscriptionAllOf) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEvalExpirationDate

`func (o *SubscriptionAllOf) GetEvalExpirationDate() time.Time`

GetEvalExpirationDate returns the EvalExpirationDate field if non-nil, zero value otherwise.

### GetEvalExpirationDateOk

`func (o *SubscriptionAllOf) GetEvalExpirationDateOk() (*time.Time, bool)`

GetEvalExpirationDateOk returns a tuple with the EvalExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalExpirationDate

`func (o *SubscriptionAllOf) SetEvalExpirationDate(v time.Time)`

SetEvalExpirationDate sets EvalExpirationDate field to given value.

### HasEvalExpirationDate

`func (o *SubscriptionAllOf) HasEvalExpirationDate() bool`

HasEvalExpirationDate returns a boolean if a field has been set.

### GetLabels

`func (o *SubscriptionAllOf) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SubscriptionAllOf) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SubscriptionAllOf) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SubscriptionAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMetrics

`func (o *SubscriptionAllOf) GetMetrics() []OneMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *SubscriptionAllOf) GetMetricsOk() (*[]OneMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *SubscriptionAllOf) SetMetrics(v []OneMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *SubscriptionAllOf) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetNotificationContacts

`func (o *SubscriptionAllOf) GetNotificationContacts() []Account`

GetNotificationContacts returns the NotificationContacts field if non-nil, zero value otherwise.

### GetNotificationContactsOk

`func (o *SubscriptionAllOf) GetNotificationContactsOk() (*[]Account, bool)`

GetNotificationContactsOk returns a tuple with the NotificationContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationContacts

`func (o *SubscriptionAllOf) SetNotificationContacts(v []Account)`

SetNotificationContacts sets NotificationContacts field to given value.

### HasNotificationContacts

`func (o *SubscriptionAllOf) HasNotificationContacts() bool`

HasNotificationContacts returns a boolean if a field has been set.

### GetPlan

`func (o *SubscriptionAllOf) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SubscriptionAllOf) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SubscriptionAllOf) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *SubscriptionAllOf) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionAllOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionAllOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionAllOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriptionAllOf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


