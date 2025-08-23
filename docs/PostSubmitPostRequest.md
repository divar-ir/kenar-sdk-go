# PostSubmitPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatEnabled** | **bool** | امکان چت فعال باشد | 
**City** | **string** | شهر آگهی | 
**Description** | **string** | توضیحات آگهی | 
**HidePhone** | **bool** | عدم نمایش شماره تماس به کاربران | 
**Images** | **[]string** |  | 
**LocationType** | [**SubmitPostRequestLocationType**](SubmitPostRequestLocationType.md) |  | 
**Title** | **string** | عنوان آگهی | 
**District** | Pointer to **string** | محله آگهی | [optional] 
**LandlineNumbers** | Pointer to **[]string** | Landline numbers to be added to the post | [optional] 
**Latitude** | Pointer to **float64** | عرض جغرافیایی آگهی | [optional] 
**Longitude** | Pointer to **float64** | طول جغرافیایی آگهی | [optional] 
**Services** | Pointer to [**OpenPlatformpostServicesFields**](OpenPlatformpostServicesFields.md) |  | [optional] 
**TemporaryResidence** | Pointer to [**PostTemporaryResidenceFields**](PostTemporaryResidenceFields.md) |  | [optional] 

## Methods

### NewPostSubmitPostRequest

`func NewPostSubmitPostRequest(chatEnabled bool, city string, description string, hidePhone bool, images []string, locationType SubmitPostRequestLocationType, title string, ) *PostSubmitPostRequest`

NewPostSubmitPostRequest instantiates a new PostSubmitPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSubmitPostRequestWithDefaults

`func NewPostSubmitPostRequestWithDefaults() *PostSubmitPostRequest`

NewPostSubmitPostRequestWithDefaults instantiates a new PostSubmitPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatEnabled

`func (o *PostSubmitPostRequest) GetChatEnabled() bool`

GetChatEnabled returns the ChatEnabled field if non-nil, zero value otherwise.

### GetChatEnabledOk

`func (o *PostSubmitPostRequest) GetChatEnabledOk() (*bool, bool)`

GetChatEnabledOk returns a tuple with the ChatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatEnabled

`func (o *PostSubmitPostRequest) SetChatEnabled(v bool)`

SetChatEnabled sets ChatEnabled field to given value.


### GetCity

`func (o *PostSubmitPostRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PostSubmitPostRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PostSubmitPostRequest) SetCity(v string)`

SetCity sets City field to given value.


### GetDescription

`func (o *PostSubmitPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostSubmitPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostSubmitPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHidePhone

`func (o *PostSubmitPostRequest) GetHidePhone() bool`

GetHidePhone returns the HidePhone field if non-nil, zero value otherwise.

### GetHidePhoneOk

`func (o *PostSubmitPostRequest) GetHidePhoneOk() (*bool, bool)`

GetHidePhoneOk returns a tuple with the HidePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePhone

`func (o *PostSubmitPostRequest) SetHidePhone(v bool)`

SetHidePhone sets HidePhone field to given value.


### GetImages

`func (o *PostSubmitPostRequest) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *PostSubmitPostRequest) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *PostSubmitPostRequest) SetImages(v []string)`

SetImages sets Images field to given value.


### GetLocationType

`func (o *PostSubmitPostRequest) GetLocationType() SubmitPostRequestLocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *PostSubmitPostRequest) GetLocationTypeOk() (*SubmitPostRequestLocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *PostSubmitPostRequest) SetLocationType(v SubmitPostRequestLocationType)`

SetLocationType sets LocationType field to given value.


### GetTitle

`func (o *PostSubmitPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostSubmitPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostSubmitPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDistrict

`func (o *PostSubmitPostRequest) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *PostSubmitPostRequest) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *PostSubmitPostRequest) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *PostSubmitPostRequest) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetLandlineNumbers

`func (o *PostSubmitPostRequest) GetLandlineNumbers() []string`

GetLandlineNumbers returns the LandlineNumbers field if non-nil, zero value otherwise.

### GetLandlineNumbersOk

`func (o *PostSubmitPostRequest) GetLandlineNumbersOk() (*[]string, bool)`

GetLandlineNumbersOk returns a tuple with the LandlineNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandlineNumbers

`func (o *PostSubmitPostRequest) SetLandlineNumbers(v []string)`

SetLandlineNumbers sets LandlineNumbers field to given value.

### HasLandlineNumbers

`func (o *PostSubmitPostRequest) HasLandlineNumbers() bool`

HasLandlineNumbers returns a boolean if a field has been set.

### GetLatitude

`func (o *PostSubmitPostRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PostSubmitPostRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PostSubmitPostRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *PostSubmitPostRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *PostSubmitPostRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PostSubmitPostRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PostSubmitPostRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *PostSubmitPostRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetServices

`func (o *PostSubmitPostRequest) GetServices() OpenPlatformpostServicesFields`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *PostSubmitPostRequest) GetServicesOk() (*OpenPlatformpostServicesFields, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *PostSubmitPostRequest) SetServices(v OpenPlatformpostServicesFields)`

SetServices sets Services field to given value.

### HasServices

`func (o *PostSubmitPostRequest) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTemporaryResidence

`func (o *PostSubmitPostRequest) GetTemporaryResidence() PostTemporaryResidenceFields`

GetTemporaryResidence returns the TemporaryResidence field if non-nil, zero value otherwise.

### GetTemporaryResidenceOk

`func (o *PostSubmitPostRequest) GetTemporaryResidenceOk() (*PostTemporaryResidenceFields, bool)`

GetTemporaryResidenceOk returns a tuple with the TemporaryResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryResidence

`func (o *PostSubmitPostRequest) SetTemporaryResidence(v PostTemporaryResidenceFields)`

SetTemporaryResidence sets TemporaryResidence field to given value.

### HasTemporaryResidence

`func (o *PostSubmitPostRequest) HasTemporaryResidence() bool`

HasTemporaryResidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


