# PostSubmitUserPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessToken** | **string** | توکن کسب‌وکاری که این آگهی متعلق به آن می‌شود | 
**CategoryFields** | **map[string]interface{}** | فیلدهای مختص دسته‌بندی که باید مطابق schema تکمیل شوند. schema را اینجا ببینید: https://kenar.divar.dev/openapi-doc/assets-get-submit-schema/ | 
**GeneralData** | [**PostPostGeneralData**](PostPostGeneralData.md) |  | 

## Methods

### NewPostSubmitUserPostRequest

`func NewPostSubmitUserPostRequest(businessToken string, categoryFields map[string]interface{}, generalData PostPostGeneralData, ) *PostSubmitUserPostRequest`

NewPostSubmitUserPostRequest instantiates a new PostSubmitUserPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSubmitUserPostRequestWithDefaults

`func NewPostSubmitUserPostRequestWithDefaults() *PostSubmitUserPostRequest`

NewPostSubmitUserPostRequestWithDefaults instantiates a new PostSubmitUserPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessToken

`func (o *PostSubmitUserPostRequest) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *PostSubmitUserPostRequest) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *PostSubmitUserPostRequest) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.


### GetCategoryFields

`func (o *PostSubmitUserPostRequest) GetCategoryFields() map[string]interface{}`

GetCategoryFields returns the CategoryFields field if non-nil, zero value otherwise.

### GetCategoryFieldsOk

`func (o *PostSubmitUserPostRequest) GetCategoryFieldsOk() (*map[string]interface{}, bool)`

GetCategoryFieldsOk returns a tuple with the CategoryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFields

`func (o *PostSubmitUserPostRequest) SetCategoryFields(v map[string]interface{})`

SetCategoryFields sets CategoryFields field to given value.


### GetGeneralData

`func (o *PostSubmitUserPostRequest) GetGeneralData() PostPostGeneralData`

GetGeneralData returns the GeneralData field if non-nil, zero value otherwise.

### GetGeneralDataOk

`func (o *PostSubmitUserPostRequest) GetGeneralDataOk() (*PostPostGeneralData, bool)`

GetGeneralDataOk returns a tuple with the GeneralData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralData

`func (o *PostSubmitUserPostRequest) SetGeneralData(v PostPostGeneralData)`

SetGeneralData sets GeneralData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


