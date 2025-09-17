# PostSubmitUserPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryFields** | **map[string]interface{}** | فیلدهای ویژه هر دسته‌بندی که باید مطابق قالب مشخص شده تکمیل شوند. قالب را از اینجا ببینید: https://divar-ir.github.io/kenar-docs/openapi-doc/assets-get-submit-schema/ | 
**GeneralData** | [**PostSubmitPostGeneralData**](PostSubmitPostGeneralData.md) |  | 

## Methods

### NewPostSubmitUserPostRequest

`func NewPostSubmitUserPostRequest(categoryFields map[string]interface{}, generalData PostSubmitPostGeneralData, ) *PostSubmitUserPostRequest`

NewPostSubmitUserPostRequest instantiates a new PostSubmitUserPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSubmitUserPostRequestWithDefaults

`func NewPostSubmitUserPostRequestWithDefaults() *PostSubmitUserPostRequest`

NewPostSubmitUserPostRequestWithDefaults instantiates a new PostSubmitUserPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *PostSubmitUserPostRequest) GetGeneralData() PostSubmitPostGeneralData`

GetGeneralData returns the GeneralData field if non-nil, zero value otherwise.

### GetGeneralDataOk

`func (o *PostSubmitUserPostRequest) GetGeneralDataOk() (*PostSubmitPostGeneralData, bool)`

GetGeneralDataOk returns a tuple with the GeneralData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralData

`func (o *PostSubmitUserPostRequest) SetGeneralData(v PostSubmitPostGeneralData)`

SetGeneralData sets GeneralData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


