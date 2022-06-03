# ConnectorNamespaceQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | Pointer to **int32** |  | [optional] 
**MemoryRequests** | Pointer to **string** | Memory quota for limits or requests | [optional] 
**MemoryLimits** | Pointer to **string** | Memory quota for limits or requests | [optional] 
**CpuRequests** | Pointer to **string** | CPU quota for limits or requests | [optional] 
**CpuLimits** | Pointer to **string** | CPU quota for limits or requests | [optional] 

## Methods

### NewConnectorNamespaceQuota

`func NewConnectorNamespaceQuota() *ConnectorNamespaceQuota`

NewConnectorNamespaceQuota instantiates a new ConnectorNamespaceQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorNamespaceQuotaWithDefaults

`func NewConnectorNamespaceQuotaWithDefaults() *ConnectorNamespaceQuota`

NewConnectorNamespaceQuotaWithDefaults instantiates a new ConnectorNamespaceQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectors

`func (o *ConnectorNamespaceQuota) GetConnectors() int32`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *ConnectorNamespaceQuota) GetConnectorsOk() (*int32, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *ConnectorNamespaceQuota) SetConnectors(v int32)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *ConnectorNamespaceQuota) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetMemoryRequests

`func (o *ConnectorNamespaceQuota) GetMemoryRequests() string`

GetMemoryRequests returns the MemoryRequests field if non-nil, zero value otherwise.

### GetMemoryRequestsOk

`func (o *ConnectorNamespaceQuota) GetMemoryRequestsOk() (*string, bool)`

GetMemoryRequestsOk returns a tuple with the MemoryRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequests

`func (o *ConnectorNamespaceQuota) SetMemoryRequests(v string)`

SetMemoryRequests sets MemoryRequests field to given value.

### HasMemoryRequests

`func (o *ConnectorNamespaceQuota) HasMemoryRequests() bool`

HasMemoryRequests returns a boolean if a field has been set.

### GetMemoryLimits

`func (o *ConnectorNamespaceQuota) GetMemoryLimits() string`

GetMemoryLimits returns the MemoryLimits field if non-nil, zero value otherwise.

### GetMemoryLimitsOk

`func (o *ConnectorNamespaceQuota) GetMemoryLimitsOk() (*string, bool)`

GetMemoryLimitsOk returns a tuple with the MemoryLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimits

`func (o *ConnectorNamespaceQuota) SetMemoryLimits(v string)`

SetMemoryLimits sets MemoryLimits field to given value.

### HasMemoryLimits

`func (o *ConnectorNamespaceQuota) HasMemoryLimits() bool`

HasMemoryLimits returns a boolean if a field has been set.

### GetCpuRequests

`func (o *ConnectorNamespaceQuota) GetCpuRequests() string`

GetCpuRequests returns the CpuRequests field if non-nil, zero value otherwise.

### GetCpuRequestsOk

`func (o *ConnectorNamespaceQuota) GetCpuRequestsOk() (*string, bool)`

GetCpuRequestsOk returns a tuple with the CpuRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequests

`func (o *ConnectorNamespaceQuota) SetCpuRequests(v string)`

SetCpuRequests sets CpuRequests field to given value.

### HasCpuRequests

`func (o *ConnectorNamespaceQuota) HasCpuRequests() bool`

HasCpuRequests returns a boolean if a field has been set.

### GetCpuLimits

`func (o *ConnectorNamespaceQuota) GetCpuLimits() string`

GetCpuLimits returns the CpuLimits field if non-nil, zero value otherwise.

### GetCpuLimitsOk

`func (o *ConnectorNamespaceQuota) GetCpuLimitsOk() (*string, bool)`

GetCpuLimitsOk returns a tuple with the CpuLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimits

`func (o *ConnectorNamespaceQuota) SetCpuLimits(v string)`

SetCpuLimits sets CpuLimits field to given value.

### HasCpuLimits

`func (o *ConnectorNamespaceQuota) HasCpuLimits() bool`

HasCpuLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


