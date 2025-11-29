# PostGetUploadURLsV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**GetUploadURLsV2ResponseUploadFormat**](GetUploadURLsV2ResponseUploadFormat.md) |  | [optional] 
**Video** | Pointer to [**GetUploadURLsV2ResponseUploadFormat**](GetUploadURLsV2ResponseUploadFormat.md) |  | [optional] 

## Methods

### NewPostGetUploadURLsV2Response

`func NewPostGetUploadURLsV2Response() *PostGetUploadURLsV2Response`

NewPostGetUploadURLsV2Response instantiates a new PostGetUploadURLsV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostGetUploadURLsV2ResponseWithDefaults

`func NewPostGetUploadURLsV2ResponseWithDefaults() *PostGetUploadURLsV2Response`

NewPostGetUploadURLsV2ResponseWithDefaults instantiates a new PostGetUploadURLsV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *PostGetUploadURLsV2Response) GetImage() GetUploadURLsV2ResponseUploadFormat`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PostGetUploadURLsV2Response) GetImageOk() (*GetUploadURLsV2ResponseUploadFormat, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PostGetUploadURLsV2Response) SetImage(v GetUploadURLsV2ResponseUploadFormat)`

SetImage sets Image field to given value.

### HasImage

`func (o *PostGetUploadURLsV2Response) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVideo

`func (o *PostGetUploadURLsV2Response) GetVideo() GetUploadURLsV2ResponseUploadFormat`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *PostGetUploadURLsV2Response) GetVideoOk() (*GetUploadURLsV2ResponseUploadFormat, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *PostGetUploadURLsV2Response) SetVideo(v GetUploadURLsV2ResponseUploadFormat)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *PostGetUploadURLsV2Response) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


