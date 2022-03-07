# InlineResponse400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string][]map[string]interface{}** |  | [optional] 
**Entity** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**EntityTag** | Pointer to [**InlineResponse400ResponseEntityTag**](InlineResponse400ResponseEntityTag.md) |  | [optional] 
**Cookies** | Pointer to [**map[string]InlineResponse400ResponseCookies**](InlineResponse400ResponseCookies.md) |  | [optional] 
**AllowedMethods** | Pointer to **[]string** |  | [optional] 
**MediaType** | Pointer to [**InlineResponse400ResponseMediaType**](InlineResponse400ResponseMediaType.md) |  | [optional] 
**StringHeaders** | Pointer to **map[string][]string** |  | [optional] 
**StatusInfo** | Pointer to [**InlineResponse400ResponseStatusInfo**](InlineResponse400ResponseStatusInfo.md) |  | [optional] 
**Links** | Pointer to [**[]InlineResponse400ResponseLinks**](InlineResponse400ResponseLinks.md) |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**Language** | Pointer to [**InlineResponse400ResponseLanguage**](InlineResponse400ResponseLanguage.md) |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**Headers** | Pointer to **map[string][]map[string]interface{}** |  | [optional] 


## Methods

### NewInlineResponse400Response

`func NewInlineResponse400Response() *InlineResponse400Response`

NewInlineResponse400Response instantiates a new InlineResponse400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse400ResponseWithDefaults

`func NewInlineResponse400ResponseWithDefaults() *InlineResponse400Response`

NewInlineResponse400ResponseWithDefaults instantiates a new InlineResponse400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetMetadata

`func (o *InlineResponse400Response) GetMetadata() map[string][]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InlineResponse400Response) GetMetadataOk() (*map[string][]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InlineResponse400Response) SetMetadata(v map[string][]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InlineResponse400Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


### GetEntity

`func (o *InlineResponse400Response) GetEntity() map[string]interface{}`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *InlineResponse400Response) GetEntityOk() (*map[string]interface{}, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *InlineResponse400Response) SetEntity(v map[string]interface{})`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *InlineResponse400Response) HasEntity() bool`

HasEntity returns a boolean if a field has been set.


### GetStatus

`func (o *InlineResponse400Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse400Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse400Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse400Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


### GetEntityTag

`func (o *InlineResponse400Response) GetEntityTag() InlineResponse400ResponseEntityTag`

GetEntityTag returns the EntityTag field if non-nil, zero value otherwise.

### GetEntityTagOk

`func (o *InlineResponse400Response) GetEntityTagOk() (*InlineResponse400ResponseEntityTag, bool)`

GetEntityTagOk returns a tuple with the EntityTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTag

`func (o *InlineResponse400Response) SetEntityTag(v InlineResponse400ResponseEntityTag)`

SetEntityTag sets EntityTag field to given value.

### HasEntityTag

`func (o *InlineResponse400Response) HasEntityTag() bool`

HasEntityTag returns a boolean if a field has been set.


### GetCookies

`func (o *InlineResponse400Response) GetCookies() map[string]InlineResponse400ResponseCookies`

GetCookies returns the Cookies field if non-nil, zero value otherwise.

### GetCookiesOk

`func (o *InlineResponse400Response) GetCookiesOk() (*map[string]InlineResponse400ResponseCookies, bool)`

GetCookiesOk returns a tuple with the Cookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookies

`func (o *InlineResponse400Response) SetCookies(v map[string]InlineResponse400ResponseCookies)`

SetCookies sets Cookies field to given value.

### HasCookies

`func (o *InlineResponse400Response) HasCookies() bool`

HasCookies returns a boolean if a field has been set.


### GetAllowedMethods

`func (o *InlineResponse400Response) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *InlineResponse400Response) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *InlineResponse400Response) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *InlineResponse400Response) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.


### GetMediaType

`func (o *InlineResponse400Response) GetMediaType() InlineResponse400ResponseMediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *InlineResponse400Response) GetMediaTypeOk() (*InlineResponse400ResponseMediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *InlineResponse400Response) SetMediaType(v InlineResponse400ResponseMediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *InlineResponse400Response) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.


### GetStringHeaders

`func (o *InlineResponse400Response) GetStringHeaders() map[string][]string`

GetStringHeaders returns the StringHeaders field if non-nil, zero value otherwise.

### GetStringHeadersOk

`func (o *InlineResponse400Response) GetStringHeadersOk() (*map[string][]string, bool)`

GetStringHeadersOk returns a tuple with the StringHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringHeaders

`func (o *InlineResponse400Response) SetStringHeaders(v map[string][]string)`

SetStringHeaders sets StringHeaders field to given value.

### HasStringHeaders

`func (o *InlineResponse400Response) HasStringHeaders() bool`

HasStringHeaders returns a boolean if a field has been set.


### GetStatusInfo

`func (o *InlineResponse400Response) GetStatusInfo() InlineResponse400ResponseStatusInfo`

GetStatusInfo returns the StatusInfo field if non-nil, zero value otherwise.

### GetStatusInfoOk

`func (o *InlineResponse400Response) GetStatusInfoOk() (*InlineResponse400ResponseStatusInfo, bool)`

GetStatusInfoOk returns a tuple with the StatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfo

`func (o *InlineResponse400Response) SetStatusInfo(v InlineResponse400ResponseStatusInfo)`

SetStatusInfo sets StatusInfo field to given value.

### HasStatusInfo

`func (o *InlineResponse400Response) HasStatusInfo() bool`

HasStatusInfo returns a boolean if a field has been set.


### GetLinks

`func (o *InlineResponse400Response) GetLinks() []InlineResponse400ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineResponse400Response) GetLinksOk() (*[]InlineResponse400ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineResponse400Response) SetLinks(v []InlineResponse400ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineResponse400Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


### GetLength

`func (o *InlineResponse400Response) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *InlineResponse400Response) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *InlineResponse400Response) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *InlineResponse400Response) HasLength() bool`

HasLength returns a boolean if a field has been set.


### GetLanguage

`func (o *InlineResponse400Response) GetLanguage() InlineResponse400ResponseLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *InlineResponse400Response) GetLanguageOk() (*InlineResponse400ResponseLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *InlineResponse400Response) SetLanguage(v InlineResponse400ResponseLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *InlineResponse400Response) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


### GetLocation

`func (o *InlineResponse400Response) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InlineResponse400Response) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InlineResponse400Response) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InlineResponse400Response) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


### GetLastModified

`func (o *InlineResponse400Response) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *InlineResponse400Response) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *InlineResponse400Response) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *InlineResponse400Response) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


### GetDate

`func (o *InlineResponse400Response) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InlineResponse400Response) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InlineResponse400Response) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *InlineResponse400Response) HasDate() bool`

HasDate returns a boolean if a field has been set.


### GetHeaders

`func (o *InlineResponse400Response) GetHeaders() map[string][]map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineResponse400Response) GetHeadersOk() (*map[string][]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineResponse400Response) SetHeaders(v map[string][]map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineResponse400Response) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

