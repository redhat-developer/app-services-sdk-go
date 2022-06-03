# UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Admin** | Pointer to **bool** |  | [optional] 
**Developer** | Pointer to **bool** |  | [optional] 
**Viewer** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserInfo

`func NewUserInfo() *UserInfo`

NewUserInfo instantiates a new UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoWithDefaults

`func NewUserInfoWithDefaults() *UserInfo`

NewUserInfoWithDefaults instantiates a new UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAdmin

`func (o *UserInfo) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserInfo) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserInfo) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *UserInfo) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetDeveloper

`func (o *UserInfo) GetDeveloper() bool`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *UserInfo) GetDeveloperOk() (*bool, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloper

`func (o *UserInfo) SetDeveloper(v bool)`

SetDeveloper sets Developer field to given value.

### HasDeveloper

`func (o *UserInfo) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### GetViewer

`func (o *UserInfo) GetViewer() bool`

GetViewer returns the Viewer field if non-nil, zero value otherwise.

### GetViewerOk

`func (o *UserInfo) GetViewerOk() (*bool, bool)`

GetViewerOk returns a tuple with the Viewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewer

`func (o *UserInfo) SetViewer(v bool)`

SetViewer sets Viewer field to given value.

### HasViewer

`func (o *UserInfo) HasViewer() bool`

HasViewer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


