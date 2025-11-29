# PostGeneralDataPostVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **string** | زمان ویدیو به ثانیه | 
**Name** | **string** | Name of the video, retrieved from &#x60;video_name&#x60; field in the response of upload video endpoint | 
**ThumbnailName** | **string** | کاور ویدیو. این مقدار را از روی فیلد &#x60;thumbnail_name&#x60; در پاسخ به درخواست آپلود ویدیو پر کنید. این تصویر، فریم اول ویدیو‌ی آپلود شده است. | 

## Methods

### NewPostGeneralDataPostVideo

`func NewPostGeneralDataPostVideo(duration string, name string, thumbnailName string, ) *PostGeneralDataPostVideo`

NewPostGeneralDataPostVideo instantiates a new PostGeneralDataPostVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostGeneralDataPostVideoWithDefaults

`func NewPostGeneralDataPostVideoWithDefaults() *PostGeneralDataPostVideo`

NewPostGeneralDataPostVideoWithDefaults instantiates a new PostGeneralDataPostVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *PostGeneralDataPostVideo) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PostGeneralDataPostVideo) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PostGeneralDataPostVideo) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetName

`func (o *PostGeneralDataPostVideo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostGeneralDataPostVideo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostGeneralDataPostVideo) SetName(v string)`

SetName sets Name field to given value.


### GetThumbnailName

`func (o *PostGeneralDataPostVideo) GetThumbnailName() string`

GetThumbnailName returns the ThumbnailName field if non-nil, zero value otherwise.

### GetThumbnailNameOk

`func (o *PostGeneralDataPostVideo) GetThumbnailNameOk() (*string, bool)`

GetThumbnailNameOk returns a tuple with the ThumbnailName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailName

`func (o *PostGeneralDataPostVideo) SetThumbnailName(v string)`

SetThumbnailName sets ThumbnailName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


