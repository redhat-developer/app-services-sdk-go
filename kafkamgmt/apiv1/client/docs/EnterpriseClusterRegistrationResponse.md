# EnterpriseClusterRegistrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** |  | 
**Href** | **string** |  | 
**ClusterId** | Pointer to **string** | OCM cluster id of the registered Enterprise cluster | [optional] 
**Status** | Pointer to **string** | status of registered Enterprise cluster | [optional] 
**FleetshardParameters** | Pointer to [**[]FleetshardParameter**](FleetshardParameter.md) |  | [optional] 

## Methods

### NewEnterpriseClusterRegistrationResponse

`func NewEnterpriseClusterRegistrationResponse(id string, kind string, href string, ) *EnterpriseClusterRegistrationResponse`

NewEnterpriseClusterRegistrationResponse instantiates a new EnterpriseClusterRegistrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseClusterRegistrationResponseWithDefaults

`func NewEnterpriseClusterRegistrationResponseWithDefaults() *EnterpriseClusterRegistrationResponse`

NewEnterpriseClusterRegistrationResponseWithDefaults instantiates a new EnterpriseClusterRegistrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnterpriseClusterRegistrationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnterpriseClusterRegistrationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnterpriseClusterRegistrationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *EnterpriseClusterRegistrationResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EnterpriseClusterRegistrationResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EnterpriseClusterRegistrationResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetHref

`func (o *EnterpriseClusterRegistrationResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *EnterpriseClusterRegistrationResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *EnterpriseClusterRegistrationResponse) SetHref(v string)`

SetHref sets Href field to given value.


### GetClusterId

`func (o *EnterpriseClusterRegistrationResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnterpriseClusterRegistrationResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnterpriseClusterRegistrationResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EnterpriseClusterRegistrationResponse) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetStatus

`func (o *EnterpriseClusterRegistrationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnterpriseClusterRegistrationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnterpriseClusterRegistrationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnterpriseClusterRegistrationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFleetshardParameters

`func (o *EnterpriseClusterRegistrationResponse) GetFleetshardParameters() []FleetshardParameter`

GetFleetshardParameters returns the FleetshardParameters field if non-nil, zero value otherwise.

### GetFleetshardParametersOk

`func (o *EnterpriseClusterRegistrationResponse) GetFleetshardParametersOk() (*[]FleetshardParameter, bool)`

GetFleetshardParametersOk returns a tuple with the FleetshardParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetshardParameters

`func (o *EnterpriseClusterRegistrationResponse) SetFleetshardParameters(v []FleetshardParameter)`

SetFleetshardParameters sets FleetshardParameters field to given value.

### HasFleetshardParameters

`func (o *EnterpriseClusterRegistrationResponse) HasFleetshardParameters() bool`

HasFleetshardParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


