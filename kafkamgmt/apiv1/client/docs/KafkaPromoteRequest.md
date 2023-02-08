# KafkaPromoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredKafkaBillingModel** | **string** | The desired Kafka billing model to promote the kafka instance to. Promotion is performed asynchronously. Accepted values: [&#39;marketplace&#39;, &#39;standard&#39;] | 
**DesiredMarketplace** | Pointer to **string** | The desired billing marketplace to promote the kafka instance to. Accepted values: [&#39;aws&#39;, &#39;rhm&#39;]. Only considered when desired_kafka_billing_model is &#39;marketplace&#39; | [optional] 
**DesiredBillingCloudAccountId** | Pointer to **string** | The desired Kafka billing cloud account ID to promote the kafka instance to. Only considered when desired_kafka_billing_model is &#39;marketplace&#39; | [optional] 

## Methods

### NewKafkaPromoteRequest

`func NewKafkaPromoteRequest(desiredKafkaBillingModel string, ) *KafkaPromoteRequest`

NewKafkaPromoteRequest instantiates a new KafkaPromoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaPromoteRequestWithDefaults

`func NewKafkaPromoteRequestWithDefaults() *KafkaPromoteRequest`

NewKafkaPromoteRequestWithDefaults instantiates a new KafkaPromoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredKafkaBillingModel

`func (o *KafkaPromoteRequest) GetDesiredKafkaBillingModel() string`

GetDesiredKafkaBillingModel returns the DesiredKafkaBillingModel field if non-nil, zero value otherwise.

### GetDesiredKafkaBillingModelOk

`func (o *KafkaPromoteRequest) GetDesiredKafkaBillingModelOk() (*string, bool)`

GetDesiredKafkaBillingModelOk returns a tuple with the DesiredKafkaBillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredKafkaBillingModel

`func (o *KafkaPromoteRequest) SetDesiredKafkaBillingModel(v string)`

SetDesiredKafkaBillingModel sets DesiredKafkaBillingModel field to given value.


### GetDesiredMarketplace

`func (o *KafkaPromoteRequest) GetDesiredMarketplace() string`

GetDesiredMarketplace returns the DesiredMarketplace field if non-nil, zero value otherwise.

### GetDesiredMarketplaceOk

`func (o *KafkaPromoteRequest) GetDesiredMarketplaceOk() (*string, bool)`

GetDesiredMarketplaceOk returns a tuple with the DesiredMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredMarketplace

`func (o *KafkaPromoteRequest) SetDesiredMarketplace(v string)`

SetDesiredMarketplace sets DesiredMarketplace field to given value.

### HasDesiredMarketplace

`func (o *KafkaPromoteRequest) HasDesiredMarketplace() bool`

HasDesiredMarketplace returns a boolean if a field has been set.

### GetDesiredBillingCloudAccountId

`func (o *KafkaPromoteRequest) GetDesiredBillingCloudAccountId() string`

GetDesiredBillingCloudAccountId returns the DesiredBillingCloudAccountId field if non-nil, zero value otherwise.

### GetDesiredBillingCloudAccountIdOk

`func (o *KafkaPromoteRequest) GetDesiredBillingCloudAccountIdOk() (*string, bool)`

GetDesiredBillingCloudAccountIdOk returns a tuple with the DesiredBillingCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredBillingCloudAccountId

`func (o *KafkaPromoteRequest) SetDesiredBillingCloudAccountId(v string)`

SetDesiredBillingCloudAccountId sets DesiredBillingCloudAccountId field to given value.

### HasDesiredBillingCloudAccountId

`func (o *KafkaPromoteRequest) HasDesiredBillingCloudAccountId() bool`

HasDesiredBillingCloudAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


