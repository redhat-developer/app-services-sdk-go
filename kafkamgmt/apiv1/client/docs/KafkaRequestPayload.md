# KafkaRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | The cloud provider where the Kafka cluster will be created in | [optional] 
**MultiAz** | Pointer to **bool** | Deprecated. If it is specified, the value will be ignored. The API automatically determines this value from the &#39;plan&#39; attribute\&quot; | [optional] 
**Name** | **string** | The name of the Kafka cluster. It must consist of lower-case alphanumeric characters or &#39;-&#39;, start with an alphabetic character, and end with an alphanumeric character, and can not be longer than 32 characters. | 
**Region** | Pointer to **string** | The region where the Kafka cluster will be created in | [optional] 
**ReauthenticationEnabled** | Pointer to **NullableBool** | Whether connection reauthentication is enabled or not. If set to true, connection reauthentication on the Kafka instance will be required every 5 minutes. The default value is true | [optional] 
**Plan** | Pointer to **string** | kafka plan in a format of &lt;instance_type&gt;.&lt;size_id&gt; | [optional] 
**BillingCloudAccountId** | Pointer to **NullableString** | cloud account id used to purchase the instance | [optional] 
**Marketplace** | Pointer to **NullableString** | marketplace where the instance is purchased on | [optional] 

## Methods

### NewKafkaRequestPayload

`func NewKafkaRequestPayload(name string, ) *KafkaRequestPayload`

NewKafkaRequestPayload instantiates a new KafkaRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaRequestPayloadWithDefaults

`func NewKafkaRequestPayloadWithDefaults() *KafkaRequestPayload`

NewKafkaRequestPayloadWithDefaults instantiates a new KafkaRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *KafkaRequestPayload) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *KafkaRequestPayload) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *KafkaRequestPayload) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *KafkaRequestPayload) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetMultiAz

`func (o *KafkaRequestPayload) GetMultiAz() bool`

GetMultiAz returns the MultiAz field if non-nil, zero value otherwise.

### GetMultiAzOk

`func (o *KafkaRequestPayload) GetMultiAzOk() (*bool, bool)`

GetMultiAzOk returns a tuple with the MultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAz

`func (o *KafkaRequestPayload) SetMultiAz(v bool)`

SetMultiAz sets MultiAz field to given value.

### HasMultiAz

`func (o *KafkaRequestPayload) HasMultiAz() bool`

HasMultiAz returns a boolean if a field has been set.

### GetName

`func (o *KafkaRequestPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaRequestPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaRequestPayload) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *KafkaRequestPayload) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *KafkaRequestPayload) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *KafkaRequestPayload) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *KafkaRequestPayload) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReauthenticationEnabled

`func (o *KafkaRequestPayload) GetReauthenticationEnabled() bool`

GetReauthenticationEnabled returns the ReauthenticationEnabled field if non-nil, zero value otherwise.

### GetReauthenticationEnabledOk

`func (o *KafkaRequestPayload) GetReauthenticationEnabledOk() (*bool, bool)`

GetReauthenticationEnabledOk returns a tuple with the ReauthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthenticationEnabled

`func (o *KafkaRequestPayload) SetReauthenticationEnabled(v bool)`

SetReauthenticationEnabled sets ReauthenticationEnabled field to given value.

### HasReauthenticationEnabled

`func (o *KafkaRequestPayload) HasReauthenticationEnabled() bool`

HasReauthenticationEnabled returns a boolean if a field has been set.

### SetReauthenticationEnabledNil

`func (o *KafkaRequestPayload) SetReauthenticationEnabledNil(b bool)`

 SetReauthenticationEnabledNil sets the value for ReauthenticationEnabled to be an explicit nil

### UnsetReauthenticationEnabled
`func (o *KafkaRequestPayload) UnsetReauthenticationEnabled()`

UnsetReauthenticationEnabled ensures that no value is present for ReauthenticationEnabled, not even an explicit nil
### GetPlan

`func (o *KafkaRequestPayload) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *KafkaRequestPayload) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *KafkaRequestPayload) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *KafkaRequestPayload) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetBillingCloudAccountId

`func (o *KafkaRequestPayload) GetBillingCloudAccountId() string`

GetBillingCloudAccountId returns the BillingCloudAccountId field if non-nil, zero value otherwise.

### GetBillingCloudAccountIdOk

`func (o *KafkaRequestPayload) GetBillingCloudAccountIdOk() (*string, bool)`

GetBillingCloudAccountIdOk returns a tuple with the BillingCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCloudAccountId

`func (o *KafkaRequestPayload) SetBillingCloudAccountId(v string)`

SetBillingCloudAccountId sets BillingCloudAccountId field to given value.

### HasBillingCloudAccountId

`func (o *KafkaRequestPayload) HasBillingCloudAccountId() bool`

HasBillingCloudAccountId returns a boolean if a field has been set.

### SetBillingCloudAccountIdNil

`func (o *KafkaRequestPayload) SetBillingCloudAccountIdNil(b bool)`

 SetBillingCloudAccountIdNil sets the value for BillingCloudAccountId to be an explicit nil

### UnsetBillingCloudAccountId
`func (o *KafkaRequestPayload) UnsetBillingCloudAccountId()`

UnsetBillingCloudAccountId ensures that no value is present for BillingCloudAccountId, not even an explicit nil
### GetMarketplace

`func (o *KafkaRequestPayload) GetMarketplace() string`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *KafkaRequestPayload) GetMarketplaceOk() (*string, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *KafkaRequestPayload) SetMarketplace(v string)`

SetMarketplace sets Marketplace field to given value.

### HasMarketplace

`func (o *KafkaRequestPayload) HasMarketplace() bool`

HasMarketplace returns a boolean if a field has been set.

### SetMarketplaceNil

`func (o *KafkaRequestPayload) SetMarketplaceNil(b bool)`

 SetMarketplaceNil sets the value for Marketplace to be an explicit nil

### UnsetMarketplace
`func (o *KafkaRequestPayload) UnsetMarketplace()`

UnsetMarketplace ensures that no value is present for Marketplace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


