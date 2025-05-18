# AddonsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenDirectLink** | Pointer to **string** | An action to send user to your URL directly with just a resource id (if applicable) | [optional] 
**OpenServerLink** | Pointer to [**AddonsOpenServerLink**](AddonsOpenServerLink.md) |  | [optional] 
**GetDynamicAction** | Pointer to [**AddonsGetDynamicAction**](AddonsGetDynamicAction.md) |  | [optional] 

## Methods

### NewAddonsAction

`func NewAddonsAction() *AddonsAction`

NewAddonsAction instantiates a new AddonsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsActionWithDefaults

`func NewAddonsActionWithDefaults() *AddonsAction`

NewAddonsActionWithDefaults instantiates a new AddonsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenDirectLink

`func (o *AddonsAction) GetOpenDirectLink() string`

GetOpenDirectLink returns the OpenDirectLink field if non-nil, zero value otherwise.

### GetOpenDirectLinkOk

`func (o *AddonsAction) GetOpenDirectLinkOk() (*string, bool)`

GetOpenDirectLinkOk returns a tuple with the OpenDirectLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDirectLink

`func (o *AddonsAction) SetOpenDirectLink(v string)`

SetOpenDirectLink sets OpenDirectLink field to given value.

### HasOpenDirectLink

`func (o *AddonsAction) HasOpenDirectLink() bool`

HasOpenDirectLink returns a boolean if a field has been set.

### GetOpenServerLink

`func (o *AddonsAction) GetOpenServerLink() AddonsOpenServerLink`

GetOpenServerLink returns the OpenServerLink field if non-nil, zero value otherwise.

### GetOpenServerLinkOk

`func (o *AddonsAction) GetOpenServerLinkOk() (*AddonsOpenServerLink, bool)`

GetOpenServerLinkOk returns a tuple with the OpenServerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenServerLink

`func (o *AddonsAction) SetOpenServerLink(v AddonsOpenServerLink)`

SetOpenServerLink sets OpenServerLink field to given value.

### HasOpenServerLink

`func (o *AddonsAction) HasOpenServerLink() bool`

HasOpenServerLink returns a boolean if a field has been set.

### GetGetDynamicAction

`func (o *AddonsAction) GetGetDynamicAction() AddonsGetDynamicAction`

GetGetDynamicAction returns the GetDynamicAction field if non-nil, zero value otherwise.

### GetGetDynamicActionOk

`func (o *AddonsAction) GetGetDynamicActionOk() (*AddonsGetDynamicAction, bool)`

GetGetDynamicActionOk returns a tuple with the GetDynamicAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetDynamicAction

`func (o *AddonsAction) SetGetDynamicAction(v AddonsGetDynamicAction)`

SetGetDynamicAction sets GetDynamicAction field to given value.

### HasGetDynamicAction

`func (o *AddonsAction) HasGetDynamicAction() bool`

HasGetDynamicAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


