# PostGetUserPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessData** | Pointer to [**PostGetUserPostResponseBusinessData**](PostGetUserPostResponseBusinessData.md) |  | [optional] 
**CategoryData** | Pointer to **map[string]interface{}** |  | [optional] 
**GeneralData** | Pointer to [**PostPostGeneralData**](PostPostGeneralData.md) |  | [optional] 
**RejectReason** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**OpenPlatformpostPostState**](OpenPlatformpostPostState.md) |  | [optional] 

## Methods

### NewPostGetUserPostResponse

`func NewPostGetUserPostResponse() *PostGetUserPostResponse`

NewPostGetUserPostResponse instantiates a new PostGetUserPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostGetUserPostResponseWithDefaults

`func NewPostGetUserPostResponseWithDefaults() *PostGetUserPostResponse`

NewPostGetUserPostResponseWithDefaults instantiates a new PostGetUserPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessData

`func (o *PostGetUserPostResponse) GetBusinessData() PostGetUserPostResponseBusinessData`

GetBusinessData returns the BusinessData field if non-nil, zero value otherwise.

### GetBusinessDataOk

`func (o *PostGetUserPostResponse) GetBusinessDataOk() (*PostGetUserPostResponseBusinessData, bool)`

GetBusinessDataOk returns a tuple with the BusinessData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessData

`func (o *PostGetUserPostResponse) SetBusinessData(v PostGetUserPostResponseBusinessData)`

SetBusinessData sets BusinessData field to given value.

### HasBusinessData

`func (o *PostGetUserPostResponse) HasBusinessData() bool`

HasBusinessData returns a boolean if a field has been set.

### GetCategoryData

`func (o *PostGetUserPostResponse) GetCategoryData() map[string]interface{}`

GetCategoryData returns the CategoryData field if non-nil, zero value otherwise.

### GetCategoryDataOk

`func (o *PostGetUserPostResponse) GetCategoryDataOk() (*map[string]interface{}, bool)`

GetCategoryDataOk returns a tuple with the CategoryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryData

`func (o *PostGetUserPostResponse) SetCategoryData(v map[string]interface{})`

SetCategoryData sets CategoryData field to given value.

### HasCategoryData

`func (o *PostGetUserPostResponse) HasCategoryData() bool`

HasCategoryData returns a boolean if a field has been set.

### GetGeneralData

`func (o *PostGetUserPostResponse) GetGeneralData() PostPostGeneralData`

GetGeneralData returns the GeneralData field if non-nil, zero value otherwise.

### GetGeneralDataOk

`func (o *PostGetUserPostResponse) GetGeneralDataOk() (*PostPostGeneralData, bool)`

GetGeneralDataOk returns a tuple with the GeneralData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralData

`func (o *PostGetUserPostResponse) SetGeneralData(v PostPostGeneralData)`

SetGeneralData sets GeneralData field to given value.

### HasGeneralData

`func (o *PostGetUserPostResponse) HasGeneralData() bool`

HasGeneralData returns a boolean if a field has been set.

### GetRejectReason

`func (o *PostGetUserPostResponse) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *PostGetUserPostResponse) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *PostGetUserPostResponse) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *PostGetUserPostResponse) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetState

`func (o *PostGetUserPostResponse) GetState() OpenPlatformpostPostState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PostGetUserPostResponse) GetStateOk() (*OpenPlatformpostPostState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PostGetUserPostResponse) SetState(v OpenPlatformpostPostState)`

SetState sets State field to given value.

### HasState

`func (o *PostGetUserPostResponse) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


