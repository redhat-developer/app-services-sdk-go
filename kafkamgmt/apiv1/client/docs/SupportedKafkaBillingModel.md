# SupportedKafkaBillingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier for the Kafka billing model | 
**AmsResource** | **string** | AMS resource to be used. Accepted values: [&#39;rhosak&#39;] | 
**AmsProduct** | **string** | AMS product to be used. Accepted values: [&#39;RHOSAK&#39;, &#39;RHOSAKTrial&#39;, &#39;RHOSAKEval&#39;, &#39;RHOSAKCC&#39;] | 
**AmsBillingModels** | **[]string** | List of AMS available billing models: Accepted values: [&#39;marketplace&#39;, &#39;marketplace-rhm&#39;, &#39;marketplace-aws&#39;] | 

## Methods

### NewSupportedKafkaBillingModel

`func NewSupportedKafkaBillingModel(id string, amsResource string, amsProduct string, amsBillingModels []string, ) *SupportedKafkaBillingModel`

NewSupportedKafkaBillingModel instantiates a new SupportedKafkaBillingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedKafkaBillingModelWithDefaults

`func NewSupportedKafkaBillingModelWithDefaults() *SupportedKafkaBillingModel`

NewSupportedKafkaBillingModelWithDefaults instantiates a new SupportedKafkaBillingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportedKafkaBillingModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportedKafkaBillingModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportedKafkaBillingModel) SetId(v string)`

SetId sets Id field to given value.


### GetAmsResource

`func (o *SupportedKafkaBillingModel) GetAmsResource() string`

GetAmsResource returns the AmsResource field if non-nil, zero value otherwise.

### GetAmsResourceOk

`func (o *SupportedKafkaBillingModel) GetAmsResourceOk() (*string, bool)`

GetAmsResourceOk returns a tuple with the AmsResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmsResource

`func (o *SupportedKafkaBillingModel) SetAmsResource(v string)`

SetAmsResource sets AmsResource field to given value.


### GetAmsProduct

`func (o *SupportedKafkaBillingModel) GetAmsProduct() string`

GetAmsProduct returns the AmsProduct field if non-nil, zero value otherwise.

### GetAmsProductOk

`func (o *SupportedKafkaBillingModel) GetAmsProductOk() (*string, bool)`

GetAmsProductOk returns a tuple with the AmsProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmsProduct

`func (o *SupportedKafkaBillingModel) SetAmsProduct(v string)`

SetAmsProduct sets AmsProduct field to given value.


### GetAmsBillingModels

`func (o *SupportedKafkaBillingModel) GetAmsBillingModels() []string`

GetAmsBillingModels returns the AmsBillingModels field if non-nil, zero value otherwise.

### GetAmsBillingModelsOk

`func (o *SupportedKafkaBillingModel) GetAmsBillingModelsOk() (*[]string, bool)`

GetAmsBillingModelsOk returns a tuple with the AmsBillingModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmsBillingModels

`func (o *SupportedKafkaBillingModel) SetAmsBillingModels(v []string)`

SetAmsBillingModels sets AmsBillingModels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


