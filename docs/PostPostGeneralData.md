# PostPostGeneralData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategorySlug** | **string** | نام دسته‌بندی هدف. دسته‌بندی‌های موجود را در این آدرس بیابید: https://kenar.divar.dev/openapi-doc/assets-get-categories/ | 
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
**Video** | Pointer to [**PostGeneralDataPostVideo**](PostGeneralDataPostVideo.md) |  | [optional] 

## Methods

### NewPostPostGeneralData

`func NewPostPostGeneralData(categorySlug string, chatEnabled bool, city string, description string, hidePhone bool, images []string, locationType PostLocationType, title string, ) *PostPostGeneralData`

NewPostPostGeneralData instantiates a new PostPostGeneralData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPostGeneralDataWithDefaults

`func NewPostPostGeneralDataWithDefaults() *PostPostGeneralData`

NewPostPostGeneralDataWithDefaults instantiates a new PostPostGeneralData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategorySlug

`func (o *PostPostGeneralData) GetCategorySlug() string`

GetCategorySlug returns the CategorySlug field if non-nil, zero value otherwise.

### GetCategorySlugOk

`func (o *PostPostGeneralData) GetCategorySlugOk() (*string, bool)`

GetCategorySlugOk returns a tuple with the CategorySlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySlug

`func (o *PostPostGeneralData) SetCategorySlug(v string)`

SetCategorySlug sets CategorySlug field to given value.


### GetChatEnabled

`func (o *PostPostGeneralData) GetChatEnabled() bool`

GetChatEnabled returns the ChatEnabled field if non-nil, zero value otherwise.

### GetChatEnabledOk

`func (o *PostPostGeneralData) GetChatEnabledOk() (*bool, bool)`

GetChatEnabledOk returns a tuple with the ChatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatEnabled

`func (o *PostPostGeneralData) SetChatEnabled(v bool)`

SetChatEnabled sets ChatEnabled field to given value.


### GetCity

`func (o *PostPostGeneralData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PostPostGeneralData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PostPostGeneralData) SetCity(v string)`

SetCity sets City field to given value.


### GetDescription

`func (o *PostPostGeneralData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostPostGeneralData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostPostGeneralData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHidePhone

`func (o *PostPostGeneralData) GetHidePhone() bool`

GetHidePhone returns the HidePhone field if non-nil, zero value otherwise.

### GetHidePhoneOk

`func (o *PostPostGeneralData) GetHidePhoneOk() (*bool, bool)`

GetHidePhoneOk returns a tuple with the HidePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePhone

`func (o *PostPostGeneralData) SetHidePhone(v bool)`

SetHidePhone sets HidePhone field to given value.


### GetImages

`func (o *PostPostGeneralData) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *PostPostGeneralData) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *PostPostGeneralData) SetImages(v []string)`

SetImages sets Images field to given value.


### GetLocationType

`func (o *PostPostGeneralData) GetLocationType() PostLocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *PostPostGeneralData) GetLocationTypeOk() (*PostLocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *PostPostGeneralData) SetLocationType(v PostLocationType)`

SetLocationType sets LocationType field to given value.


### GetTitle

`func (o *PostPostGeneralData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostPostGeneralData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostPostGeneralData) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDistrict

`func (o *PostPostGeneralData) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *PostPostGeneralData) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *PostPostGeneralData) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *PostPostGeneralData) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetLatitude

`func (o *PostPostGeneralData) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PostPostGeneralData) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PostPostGeneralData) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PostPostGeneralData) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *PostPostGeneralData) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PostPostGeneralData) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PostPostGeneralData) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PostPostGeneralData) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetVideo

`func (o *PostPostGeneralData) GetVideo() PostGeneralDataPostVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *PostPostGeneralData) GetVideoOk() (*PostGeneralDataPostVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *PostPostGeneralData) SetVideo(v PostGeneralDataPostVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *PostPostGeneralData) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


