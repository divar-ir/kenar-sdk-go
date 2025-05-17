# AddonsAddonMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**App** | Pointer to [**AppsApp**](AppsApp.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**AddonsStatus**](AddonsStatus.md) |  | [optional] 
**ServiceTags** | Pointer to [**[]AppsServiceTag**](AppsServiceTag.md) |  | [optional] 

## Methods

### NewAddonsAddonMetaData

`func NewAddonsAddonMetaData() *AddonsAddonMetaData`

NewAddonsAddonMetaData instantiates a new AddonsAddonMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsAddonMetaDataWithDefaults

`func NewAddonsAddonMetaDataWithDefaults() *AddonsAddonMetaData`

NewAddonsAddonMetaDataWithDefaults instantiates a new AddonsAddonMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddonsAddonMetaData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonsAddonMetaData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonsAddonMetaData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddonsAddonMetaData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApp

`func (o *AddonsAddonMetaData) GetApp() AppsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AddonsAddonMetaData) GetAppOk() (*AppsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AddonsAddonMetaData) SetApp(v AppsApp)`

SetApp sets App field to given value.

### HasApp

`func (o *AddonsAddonMetaData) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AddonsAddonMetaData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AddonsAddonMetaData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AddonsAddonMetaData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AddonsAddonMetaData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastModified

`func (o *AddonsAddonMetaData) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AddonsAddonMetaData) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AddonsAddonMetaData) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AddonsAddonMetaData) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetStatus

`func (o *AddonsAddonMetaData) GetStatus() AddonsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddonsAddonMetaData) GetStatusOk() (*AddonsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddonsAddonMetaData) SetStatus(v AddonsStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddonsAddonMetaData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetServiceTags

`func (o *AddonsAddonMetaData) GetServiceTags() []AppsServiceTag`

GetServiceTags returns the ServiceTags field if non-nil, zero value otherwise.

### GetServiceTagsOk

`func (o *AddonsAddonMetaData) GetServiceTagsOk() (*[]AppsServiceTag, bool)`

GetServiceTagsOk returns a tuple with the ServiceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTags

`func (o *AddonsAddonMetaData) SetServiceTags(v []AppsServiceTag)`

SetServiceTags sets ServiceTags field to given value.

### HasServiceTags

`func (o *AddonsAddonMetaData) HasServiceTags() bool`

HasServiceTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


