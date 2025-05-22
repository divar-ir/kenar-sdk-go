# AddonsBusinessAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessRef** | Pointer to **string** |  | [optional] 
**MetaData** | Pointer to [**AddonsAddonMetaData**](AddonsAddonMetaData.md) |  | [optional] 
**Widgets** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAddonsBusinessAddon

`func NewAddonsBusinessAddon() *AddonsBusinessAddon`

NewAddonsBusinessAddon instantiates a new AddonsBusinessAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsBusinessAddonWithDefaults

`func NewAddonsBusinessAddonWithDefaults() *AddonsBusinessAddon`

NewAddonsBusinessAddonWithDefaults instantiates a new AddonsBusinessAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessRef

`func (o *AddonsBusinessAddon) GetBusinessRef() string`

GetBusinessRef returns the BusinessRef field if non-nil, zero value otherwise.

### GetBusinessRefOk

`func (o *AddonsBusinessAddon) GetBusinessRefOk() (*string, bool)`

GetBusinessRefOk returns a tuple with the BusinessRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessRef

`func (o *AddonsBusinessAddon) SetBusinessRef(v string)`

SetBusinessRef sets BusinessRef field to given value.

### HasBusinessRef

`func (o *AddonsBusinessAddon) HasBusinessRef() bool`

HasBusinessRef returns a boolean if a field has been set.

### GetMetaData

`func (o *AddonsBusinessAddon) GetMetaData() AddonsAddonMetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AddonsBusinessAddon) GetMetaDataOk() (*AddonsAddonMetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AddonsBusinessAddon) SetMetaData(v AddonsAddonMetaData)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AddonsBusinessAddon) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetWidgets

`func (o *AddonsBusinessAddon) GetWidgets() map[string]interface{}`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *AddonsBusinessAddon) GetWidgetsOk() (*map[string]interface{}, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *AddonsBusinessAddon) SetWidgets(v map[string]interface{})`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *AddonsBusinessAddon) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


