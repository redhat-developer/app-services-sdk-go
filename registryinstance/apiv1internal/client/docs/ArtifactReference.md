# ArtifactReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** |  | 
**ArtifactId** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewArtifactReference

`func NewArtifactReference(groupId string, artifactId string, name string, ) *ArtifactReference`

NewArtifactReference instantiates a new ArtifactReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactReferenceWithDefaults

`func NewArtifactReferenceWithDefaults() *ArtifactReference`

NewArtifactReferenceWithDefaults instantiates a new ArtifactReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ArtifactReference) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ArtifactReference) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ArtifactReference) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetArtifactId

`func (o *ArtifactReference) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *ArtifactReference) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *ArtifactReference) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.


### GetVersion

`func (o *ArtifactReference) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArtifactReference) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArtifactReference) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ArtifactReference) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *ArtifactReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactReference) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


