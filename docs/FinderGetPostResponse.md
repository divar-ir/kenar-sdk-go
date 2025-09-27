# FinderGetPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessData** | Pointer to [**FinderGetPostResponseBusinessData**](FinderGetPostResponseBusinessData.md) |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ChatEnabled** | Pointer to **bool** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**District** | Pointer to **string** |  | [optional] 
**FirstPublishedAt** | Pointer to **time.Time** |  | [optional] 
**IsPhoneHidden** | Pointer to **bool** |  | [optional] 
**LastModifiedAt** | Pointer to **time.Time** |  | [optional] 
**State** | Pointer to [**FinderPostExtState**](FinderPostExtState.md) |  | [optional] 
**SupplierChatAssistantEnabled** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewFinderGetPostResponse

`func NewFinderGetPostResponse() *FinderGetPostResponse`

NewFinderGetPostResponse instantiates a new FinderGetPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinderGetPostResponseWithDefaults

`func NewFinderGetPostResponseWithDefaults() *FinderGetPostResponse`

NewFinderGetPostResponseWithDefaults instantiates a new FinderGetPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessData

`func (o *FinderGetPostResponse) GetBusinessData() FinderGetPostResponseBusinessData`

GetBusinessData returns the BusinessData field if non-nil, zero value otherwise.

### GetBusinessDataOk

`func (o *FinderGetPostResponse) GetBusinessDataOk() (*FinderGetPostResponseBusinessData, bool)`

GetBusinessDataOk returns a tuple with the BusinessData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessData

`func (o *FinderGetPostResponse) SetBusinessData(v FinderGetPostResponseBusinessData)`

SetBusinessData sets BusinessData field to given value.

### HasBusinessData

`func (o *FinderGetPostResponse) HasBusinessData() bool`

HasBusinessData returns a boolean if a field has been set.

### GetCategory

`func (o *FinderGetPostResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FinderGetPostResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FinderGetPostResponse) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FinderGetPostResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetChatEnabled

`func (o *FinderGetPostResponse) GetChatEnabled() bool`

GetChatEnabled returns the ChatEnabled field if non-nil, zero value otherwise.

### GetChatEnabledOk

`func (o *FinderGetPostResponse) GetChatEnabledOk() (*bool, bool)`

GetChatEnabledOk returns a tuple with the ChatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatEnabled

`func (o *FinderGetPostResponse) SetChatEnabled(v bool)`

SetChatEnabled sets ChatEnabled field to given value.

### HasChatEnabled

`func (o *FinderGetPostResponse) HasChatEnabled() bool`

HasChatEnabled returns a boolean if a field has been set.

### GetCity

`func (o *FinderGetPostResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *FinderGetPostResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *FinderGetPostResponse) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *FinderGetPostResponse) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetData

`func (o *FinderGetPostResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FinderGetPostResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FinderGetPostResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *FinderGetPostResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDistrict

`func (o *FinderGetPostResponse) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *FinderGetPostResponse) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *FinderGetPostResponse) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *FinderGetPostResponse) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetFirstPublishedAt

`func (o *FinderGetPostResponse) GetFirstPublishedAt() time.Time`

GetFirstPublishedAt returns the FirstPublishedAt field if non-nil, zero value otherwise.

### GetFirstPublishedAtOk

`func (o *FinderGetPostResponse) GetFirstPublishedAtOk() (*time.Time, bool)`

GetFirstPublishedAtOk returns a tuple with the FirstPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPublishedAt

`func (o *FinderGetPostResponse) SetFirstPublishedAt(v time.Time)`

SetFirstPublishedAt sets FirstPublishedAt field to given value.

### HasFirstPublishedAt

`func (o *FinderGetPostResponse) HasFirstPublishedAt() bool`

HasFirstPublishedAt returns a boolean if a field has been set.

### GetIsPhoneHidden

`func (o *FinderGetPostResponse) GetIsPhoneHidden() bool`

GetIsPhoneHidden returns the IsPhoneHidden field if non-nil, zero value otherwise.

### GetIsPhoneHiddenOk

`func (o *FinderGetPostResponse) GetIsPhoneHiddenOk() (*bool, bool)`

GetIsPhoneHiddenOk returns a tuple with the IsPhoneHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPhoneHidden

`func (o *FinderGetPostResponse) SetIsPhoneHidden(v bool)`

SetIsPhoneHidden sets IsPhoneHidden field to given value.

### HasIsPhoneHidden

`func (o *FinderGetPostResponse) HasIsPhoneHidden() bool`

HasIsPhoneHidden returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *FinderGetPostResponse) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *FinderGetPostResponse) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *FinderGetPostResponse) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *FinderGetPostResponse) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetState

`func (o *FinderGetPostResponse) GetState() FinderPostExtState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FinderGetPostResponse) GetStateOk() (*FinderPostExtState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FinderGetPostResponse) SetState(v FinderPostExtState)`

SetState sets State field to given value.

### HasState

`func (o *FinderGetPostResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSupplierChatAssistantEnabled

`func (o *FinderGetPostResponse) GetSupplierChatAssistantEnabled() bool`

GetSupplierChatAssistantEnabled returns the SupplierChatAssistantEnabled field if non-nil, zero value otherwise.

### GetSupplierChatAssistantEnabledOk

`func (o *FinderGetPostResponse) GetSupplierChatAssistantEnabledOk() (*bool, bool)`

GetSupplierChatAssistantEnabledOk returns a tuple with the SupplierChatAssistantEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierChatAssistantEnabled

`func (o *FinderGetPostResponse) SetSupplierChatAssistantEnabled(v bool)`

SetSupplierChatAssistantEnabled sets SupplierChatAssistantEnabled field to given value.

### HasSupplierChatAssistantEnabled

`func (o *FinderGetPostResponse) HasSupplierChatAssistantEnabled() bool`

HasSupplierChatAssistantEnabled returns a boolean if a field has been set.

### GetToken

`func (o *FinderGetPostResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FinderGetPostResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FinderGetPostResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FinderGetPostResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


