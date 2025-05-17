# AddonsGetUserAddonsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]AddonsUserAddon**](AddonsUserAddon.md) |  | [optional] 

## Methods

### NewAddonsGetUserAddonsResponse

`func NewAddonsGetUserAddonsResponse() *AddonsGetUserAddonsResponse`

NewAddonsGetUserAddonsResponse instantiates a new AddonsGetUserAddonsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsGetUserAddonsResponseWithDefaults

`func NewAddonsGetUserAddonsResponseWithDefaults() *AddonsGetUserAddonsResponse`

NewAddonsGetUserAddonsResponseWithDefaults instantiates a new AddonsGetUserAddonsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *AddonsGetUserAddonsResponse) GetAddons() []AddonsUserAddon`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *AddonsGetUserAddonsResponse) GetAddonsOk() (*[]AddonsUserAddon, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *AddonsGetUserAddonsResponse) SetAddons(v []AddonsUserAddon)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *AddonsGetUserAddonsResponse) HasAddons() bool`

HasAddons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


