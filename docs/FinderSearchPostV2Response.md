# FinderSearchPostV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Posts** | Pointer to [**[]FinderSearchPostItem**](FinderSearchPostItem.md) |  | [optional] 

## Methods

### NewFinderSearchPostV2Response

`func NewFinderSearchPostV2Response() *FinderSearchPostV2Response`

NewFinderSearchPostV2Response instantiates a new FinderSearchPostV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinderSearchPostV2ResponseWithDefaults

`func NewFinderSearchPostV2ResponseWithDefaults() *FinderSearchPostV2Response`

NewFinderSearchPostV2ResponseWithDefaults instantiates a new FinderSearchPostV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosts

`func (o *FinderSearchPostV2Response) GetPosts() []FinderSearchPostItem`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *FinderSearchPostV2Response) GetPostsOk() (*[]FinderSearchPostItem, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *FinderSearchPostV2Response) SetPosts(v []FinderSearchPostItem)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *FinderSearchPostV2Response) HasPosts() bool`

HasPosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


