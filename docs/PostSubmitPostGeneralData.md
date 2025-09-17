# PostSubmitPostGeneralData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategorySlug** | **string** | نام دسته‌بندی هدف. دسته‌بندی‌های موجود را در این آدرس بیابید: https://divar-ir.github.io/kenar-docs/openapi-doc/assets-get-categories/ | 
**ChatEnabled** | **bool** | امکان چت فعال باشد | 
**City** | **string** | شهر آگهی | 
**Description** | **string** | توضیحات آگهی | 
**HidePhone** | **bool** | عدم نمایش شماره تماس به کاربران | 
**Images** | **[]string** |  | 
**LocationType** | [**PostLocationType**](PostLocationType.md) |  | 
**Title** | **string** | عنوان آگهی | 
**District** | Pointer to **string** | محله آگهی | [optional] 
**Latitude** | Pointer to **float64** | عرض جغرافیایی آگهی | [optional] 
**Longitude** | Pointer to **float64** | طول جغرافیایی آگهی | [optional] 

## Methods

### NewPostSubmitPostGeneralData

`func NewPostSubmitPostGeneralData(categorySlug string, chatEnabled bool, city string, description string, hidePhone bool, images []string, locationType PostLocationType, title string, ) *PostSubmitPostGeneralData`

NewPostSubmitPostGeneralData instantiates a new PostSubmitPostGeneralData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSubmitPostGeneralDataWithDefaults

`func NewPostSubmitPostGeneralDataWithDefaults() *PostSubmitPostGeneralData`

NewPostSubmitPostGeneralDataWithDefaults instantiates a new PostSubmitPostGeneralData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategorySlug

`func (o *PostSubmitPostGeneralData) GetCategorySlug() string`

GetCategorySlug returns the CategorySlug field if non-nil, zero value otherwise.

### GetCategorySlugOk

`func (o *PostSubmitPostGeneralData) GetCategorySlugOk() (*string, bool)`

GetCategorySlugOk returns a tuple with the CategorySlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySlug

`func (o *PostSubmitPostGeneralData) SetCategorySlug(v string)`

SetCategorySlug sets CategorySlug field to given value.


### GetChatEnabled

`func (o *PostSubmitPostGeneralData) GetChatEnabled() bool`

GetChatEnabled returns the ChatEnabled field if non-nil, zero value otherwise.

### GetChatEnabledOk

`func (o *PostSubmitPostGeneralData) GetChatEnabledOk() (*bool, bool)`

GetChatEnabledOk returns a tuple with the ChatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatEnabled

`func (o *PostSubmitPostGeneralData) SetChatEnabled(v bool)`

SetChatEnabled sets ChatEnabled field to given value.


### GetCity

`func (o *PostSubmitPostGeneralData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PostSubmitPostGeneralData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PostSubmitPostGeneralData) SetCity(v string)`

SetCity sets City field to given value.


### GetDescription

`func (o *PostSubmitPostGeneralData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostSubmitPostGeneralData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostSubmitPostGeneralData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHidePhone

`func (o *PostSubmitPostGeneralData) GetHidePhone() bool`

GetHidePhone returns the HidePhone field if non-nil, zero value otherwise.

### GetHidePhoneOk

`func (o *PostSubmitPostGeneralData) GetHidePhoneOk() (*bool, bool)`

GetHidePhoneOk returns a tuple with the HidePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePhone

`func (o *PostSubmitPostGeneralData) SetHidePhone(v bool)`

SetHidePhone sets HidePhone field to given value.


### GetImages

`func (o *PostSubmitPostGeneralData) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *PostSubmitPostGeneralData) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *PostSubmitPostGeneralData) SetImages(v []string)`

SetImages sets Images field to given value.


### GetLocationType

`func (o *PostSubmitPostGeneralData) GetLocationType() PostLocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *PostSubmitPostGeneralData) GetLocationTypeOk() (*PostLocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *PostSubmitPostGeneralData) SetLocationType(v PostLocationType)`

SetLocationType sets LocationType field to given value.


### GetTitle

`func (o *PostSubmitPostGeneralData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostSubmitPostGeneralData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostSubmitPostGeneralData) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDistrict

`func (o *PostSubmitPostGeneralData) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *PostSubmitPostGeneralData) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *PostSubmitPostGeneralData) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *PostSubmitPostGeneralData) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetLatitude

`func (o *PostSubmitPostGeneralData) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PostSubmitPostGeneralData) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PostSubmitPostGeneralData) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PostSubmitPostGeneralData) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *PostSubmitPostGeneralData) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PostSubmitPostGeneralData) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PostSubmitPostGeneralData) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PostSubmitPostGeneralData) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


