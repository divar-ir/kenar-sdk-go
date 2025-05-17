# ManagementDevelopmentPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preset** | Pointer to [**ManagementPreset**](ManagementPreset.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**MngToken** | Pointer to **string** |  | [optional] 

## Methods

### NewManagementDevelopmentPost

`func NewManagementDevelopmentPost() *ManagementDevelopmentPost`

NewManagementDevelopmentPost instantiates a new ManagementDevelopmentPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementDevelopmentPostWithDefaults

`func NewManagementDevelopmentPostWithDefaults() *ManagementDevelopmentPost`

NewManagementDevelopmentPostWithDefaults instantiates a new ManagementDevelopmentPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreset

`func (o *ManagementDevelopmentPost) GetPreset() ManagementPreset`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *ManagementDevelopmentPost) GetPresetOk() (*ManagementPreset, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *ManagementDevelopmentPost) SetPreset(v ManagementPreset)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *ManagementDevelopmentPost) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ManagementDevelopmentPost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManagementDevelopmentPost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManagementDevelopmentPost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ManagementDevelopmentPost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetToken

`func (o *ManagementDevelopmentPost) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ManagementDevelopmentPost) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ManagementDevelopmentPost) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ManagementDevelopmentPost) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetMngToken

`func (o *ManagementDevelopmentPost) GetMngToken() string`

GetMngToken returns the MngToken field if non-nil, zero value otherwise.

### GetMngTokenOk

`func (o *ManagementDevelopmentPost) GetMngTokenOk() (*string, bool)`

GetMngTokenOk returns a tuple with the MngToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMngToken

`func (o *ManagementDevelopmentPost) SetMngToken(v string)`

SetMngToken sets MngToken field to given value.

### HasMngToken

`func (o *ManagementDevelopmentPost) HasMngToken() bool`

HasMngToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


