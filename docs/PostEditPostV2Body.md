# PostEditPostV2Body

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateMask** | **[]string** | مشخص می‌کند کدام فیلدها به‌روزرسانی شوند. از مسیرهای تودرتو مانند &#39;general_data.title&#39; یا &#39;category_data.price&#39; استفاده کنید. این کار تمایز بین پاک کردن یک فیلد و تغییر ندادن آن را ممکن می‌سازد. | 
**CategoryData** | Pointer to **map[string]interface{}** | فیلدهای مختص دسته‌بندی که باید مطابق schema تکمیل شوند. schema را اینجا ببینید: https://kenar.divar.dev/openapi-doc/assets-get-submit-schema/ | [optional] 
**GeneralData** | Pointer to [**PostPostGeneralData**](PostPostGeneralData.md) |  | [optional] 

## Methods

### NewPostEditPostV2Body

`func NewPostEditPostV2Body(updateMask []string, ) *PostEditPostV2Body`

NewPostEditPostV2Body instantiates a new PostEditPostV2Body object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEditPostV2BodyWithDefaults

`func NewPostEditPostV2BodyWithDefaults() *PostEditPostV2Body`

NewPostEditPostV2BodyWithDefaults instantiates a new PostEditPostV2Body object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateMask

`func (o *PostEditPostV2Body) GetUpdateMask() []string`

GetUpdateMask returns the UpdateMask field if non-nil, zero value otherwise.

### GetUpdateMaskOk

`func (o *PostEditPostV2Body) GetUpdateMaskOk() (*[]string, bool)`

GetUpdateMaskOk returns a tuple with the UpdateMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMask

`func (o *PostEditPostV2Body) SetUpdateMask(v []string)`

SetUpdateMask sets UpdateMask field to given value.


### GetCategoryData

`func (o *PostEditPostV2Body) GetCategoryData() map[string]interface{}`

GetCategoryData returns the CategoryData field if non-nil, zero value otherwise.

### GetCategoryDataOk

`func (o *PostEditPostV2Body) GetCategoryDataOk() (*map[string]interface{}, bool)`

GetCategoryDataOk returns a tuple with the CategoryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryData

`func (o *PostEditPostV2Body) SetCategoryData(v map[string]interface{})`

SetCategoryData sets CategoryData field to given value.

### HasCategoryData

`func (o *PostEditPostV2Body) HasCategoryData() bool`

HasCategoryData returns a boolean if a field has been set.

### GetGeneralData

`func (o *PostEditPostV2Body) GetGeneralData() PostPostGeneralData`

GetGeneralData returns the GeneralData field if non-nil, zero value otherwise.

### GetGeneralDataOk

`func (o *PostEditPostV2Body) GetGeneralDataOk() (*PostPostGeneralData, bool)`

GetGeneralDataOk returns a tuple with the GeneralData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralData

`func (o *PostEditPostV2Body) SetGeneralData(v PostPostGeneralData)`

SetGeneralData sets GeneralData field to given value.

### HasGeneralData

`func (o *PostEditPostV2Body) HasGeneralData() bool`

HasGeneralData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


