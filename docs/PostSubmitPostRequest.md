# PostSubmitPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatEnabled** | Pointer to **bool** | Whether to enable chat | [optional] 
**City** | Pointer to **string** | City of the post | [optional] 
**Description** | Pointer to **string** | Description of the post | [optional] 
**District** | Pointer to **string** | District of the post | [optional] 
**HidePhone** | Pointer to **bool** | Whether to hide the phone number from demand users | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 
**Latitude** | Pointer to **float64** | Latitude of the post | [optional] 
**Longitude** | Pointer to **float64** | Longitude of the post | [optional] 
**TemporaryResidence** | Pointer to [**PostTemporaryResidenceFields**](PostTemporaryResidenceFields.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the post | [optional] 

## Methods

### NewPostSubmitPostRequest

`func NewPostSubmitPostRequest() *PostSubmitPostRequest`

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

### HasChatEnabled

`func (o *PostSubmitPostRequest) HasChatEnabled() bool`

HasChatEnabled returns a boolean if a field has been set.

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

### HasCity

`func (o *PostSubmitPostRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

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

### HasDescription

`func (o *PostSubmitPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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

### HasHidePhone

`func (o *PostSubmitPostRequest) HasHidePhone() bool`

HasHidePhone returns a boolean if a field has been set.

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

### HasImages

`func (o *PostSubmitPostRequest) HasImages() bool`

HasImages returns a boolean if a field has been set.

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

### HasTitle

`func (o *PostSubmitPostRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


