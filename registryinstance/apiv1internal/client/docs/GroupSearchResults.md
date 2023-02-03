# GroupSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]SearchedGroup**](SearchedGroup.md) | The groups returned in the result set. | 
**Count** | **int32** | The total number of groups that matched the query that produced the result set (may be  more than the number of groups in the result set). | 

## Methods

### NewGroupSearchResults

`func NewGroupSearchResults(groups []SearchedGroup, count int32, ) *GroupSearchResults`

NewGroupSearchResults instantiates a new GroupSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSearchResultsWithDefaults

`func NewGroupSearchResultsWithDefaults() *GroupSearchResults`

NewGroupSearchResultsWithDefaults instantiates a new GroupSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *GroupSearchResults) GetGroups() []SearchedGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupSearchResults) GetGroupsOk() (*[]SearchedGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupSearchResults) SetGroups(v []SearchedGroup)`

SetGroups sets Groups field to given value.


### GetCount

`func (o *GroupSearchResults) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GroupSearchResults) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GroupSearchResults) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


