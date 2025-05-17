# AddonsButtonBar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Action** | Pointer to [**AddonsAction**](AddonsAction.md) |  | [optional] 

## Methods

### NewAddonsButtonBar

`func NewAddonsButtonBar() *AddonsButtonBar`

NewAddonsButtonBar instantiates a new AddonsButtonBar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsButtonBarWithDefaults

`func NewAddonsButtonBarWithDefaults() *AddonsButtonBar`

NewAddonsButtonBarWithDefaults instantiates a new AddonsButtonBar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AddonsButtonBar) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AddonsButtonBar) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AddonsButtonBar) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AddonsButtonBar) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAction

`func (o *AddonsButtonBar) GetAction() AddonsAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AddonsButtonBar) GetActionOk() (*AddonsAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AddonsButtonBar) SetAction(v AddonsAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *AddonsButtonBar) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


