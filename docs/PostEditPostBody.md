# PostEditPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ImagePaths** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPostEditPostBody

`func NewPostEditPostBody() *PostEditPostBody`

NewPostEditPostBody instantiates a new PostEditPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostEditPostBodyWithDefaults

`func NewPostEditPostBodyWithDefaults() *PostEditPostBody`

NewPostEditPostBodyWithDefaults instantiates a new PostEditPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PostEditPostBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostEditPostBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostEditPostBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PostEditPostBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *PostEditPostBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostEditPostBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostEditPostBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostEditPostBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImagePaths

`func (o *PostEditPostBody) GetImagePaths() []string`

GetImagePaths returns the ImagePaths field if non-nil, zero value otherwise.

### GetImagePathsOk

`func (o *PostEditPostBody) GetImagePathsOk() (*[]string, bool)`

GetImagePathsOk returns a tuple with the ImagePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePaths

`func (o *PostEditPostBody) SetImagePaths(v []string)`

SetImagePaths sets ImagePaths field to given value.

### HasImagePaths

`func (o *PostEditPostBody) HasImagePaths() bool`

HasImagePaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


