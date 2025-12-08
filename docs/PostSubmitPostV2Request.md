# PostSubmitPostV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessToken** | **string** | توکن کسب‌وکاری که این آگهی متعلق به آن می‌شود | 
**CategoryFields** | **map[string]interface{}** | فیلدهای مختص دسته‌بندی که باید مطابق schema تکمیل شوند. schema را اینجا ببینید: https://kenar.divar.dev/openapi-doc/assets-get-submit-schema/ | 
**GeneralData** | [**PostPostGeneralData**](PostPostGeneralData.md) |  | 
**LandlineNumbers** | Pointer to **[]string** | شماره‌های ثابت برای افزودن به آگهی | [optional] 

## Methods

### NewPostSubmitPostV2Request

`func NewPostSubmitPostV2Request(businessToken string, categoryFields map[string]interface{}, generalData PostPostGeneralData, ) *PostSubmitPostV2Request`

NewPostSubmitPostV2Request instantiates a new PostSubmitPostV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSubmitPostV2RequestWithDefaults

`func NewPostSubmitPostV2RequestWithDefaults() *PostSubmitPostV2Request`

NewPostSubmitPostV2RequestWithDefaults instantiates a new PostSubmitPostV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessToken

`func (o *PostSubmitPostV2Request) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *PostSubmitPostV2Request) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *PostSubmitPostV2Request) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.


### GetCategoryFields

`func (o *PostSubmitPostV2Request) GetCategoryFields() map[string]interface{}`

GetCategoryFields returns the CategoryFields field if non-nil, zero value otherwise.

### GetCategoryFieldsOk

`func (o *PostSubmitPostV2Request) GetCategoryFieldsOk() (*map[string]interface{}, bool)`

GetCategoryFieldsOk returns a tuple with the CategoryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFields

`func (o *PostSubmitPostV2Request) SetCategoryFields(v map[string]interface{})`

SetCategoryFields sets CategoryFields field to given value.


### GetGeneralData

`func (o *PostSubmitPostV2Request) GetGeneralData() PostPostGeneralData`

GetGeneralData returns the GeneralData field if non-nil, zero value otherwise.

### GetGeneralDataOk

`func (o *PostSubmitPostV2Request) GetGeneralDataOk() (*PostPostGeneralData, bool)`

GetGeneralDataOk returns a tuple with the GeneralData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralData

`func (o *PostSubmitPostV2Request) SetGeneralData(v PostPostGeneralData)`

SetGeneralData sets GeneralData field to given value.


### GetLandlineNumbers

`func (o *PostSubmitPostV2Request) GetLandlineNumbers() []string`

GetLandlineNumbers returns the LandlineNumbers field if non-nil, zero value otherwise.

### GetLandlineNumbersOk

`func (o *PostSubmitPostV2Request) GetLandlineNumbersOk() (*[]string, bool)`

GetLandlineNumbersOk returns a tuple with the LandlineNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandlineNumbers

`func (o *PostSubmitPostV2Request) SetLandlineNumbers(v []string)`

SetLandlineNumbers sets LandlineNumbers field to given value.

### HasLandlineNumbers

`func (o *PostSubmitPostV2Request) HasLandlineNumbers() bool`

HasLandlineNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


