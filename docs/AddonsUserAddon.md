# AddonsUserAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DivarUserId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**AddonsUserAddonFilters**](AddonsUserAddonFilters.md) |  | [optional] 
**MetaData** | Pointer to [**AddonsAddonMetaData**](AddonsAddonMetaData.md) |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Semantic** | Pointer to **map[string]string** |  | [optional] 
**SemanticData** | Pointer to [**AddonsAddonSemantic**](AddonsAddonSemantic.md) |  | [optional] 
**SensitiveSemantic** | Pointer to **map[string]string** |  | [optional] 
**Widgets** | Pointer to **map[string]interface{}** |  | [optional] 
**WidgetsSemantic** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAddonsUserAddon

`func NewAddonsUserAddon() *AddonsUserAddon`

NewAddonsUserAddon instantiates a new AddonsUserAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsUserAddonWithDefaults

`func NewAddonsUserAddonWithDefaults() *AddonsUserAddon`

NewAddonsUserAddonWithDefaults instantiates a new AddonsUserAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivarUserId

`func (o *AddonsUserAddon) GetDivarUserId() string`

GetDivarUserId returns the DivarUserId field if non-nil, zero value otherwise.

### GetDivarUserIdOk

`func (o *AddonsUserAddon) GetDivarUserIdOk() (*string, bool)`

GetDivarUserIdOk returns a tuple with the DivarUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivarUserId

`func (o *AddonsUserAddon) SetDivarUserId(v string)`

SetDivarUserId sets DivarUserId field to given value.

### HasDivarUserId

`func (o *AddonsUserAddon) HasDivarUserId() bool`

HasDivarUserId returns a boolean if a field has been set.

### GetFilters

`func (o *AddonsUserAddon) GetFilters() AddonsUserAddonFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AddonsUserAddon) GetFiltersOk() (*AddonsUserAddonFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AddonsUserAddon) SetFilters(v AddonsUserAddonFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AddonsUserAddon) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMetaData

`func (o *AddonsUserAddon) GetMetaData() AddonsAddonMetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AddonsUserAddon) GetMetaDataOk() (*AddonsAddonMetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AddonsUserAddon) SetMetaData(v AddonsAddonMetaData)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AddonsUserAddon) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetPhone

`func (o *AddonsUserAddon) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddonsUserAddon) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddonsUserAddon) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddonsUserAddon) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetSemantic

`func (o *AddonsUserAddon) GetSemantic() map[string]string`

GetSemantic returns the Semantic field if non-nil, zero value otherwise.

### GetSemanticOk

`func (o *AddonsUserAddon) GetSemanticOk() (*map[string]string, bool)`

GetSemanticOk returns a tuple with the Semantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemantic

`func (o *AddonsUserAddon) SetSemantic(v map[string]string)`

SetSemantic sets Semantic field to given value.

### HasSemantic

`func (o *AddonsUserAddon) HasSemantic() bool`

HasSemantic returns a boolean if a field has been set.

### GetSemanticData

`func (o *AddonsUserAddon) GetSemanticData() AddonsAddonSemantic`

GetSemanticData returns the SemanticData field if non-nil, zero value otherwise.

### GetSemanticDataOk

`func (o *AddonsUserAddon) GetSemanticDataOk() (*AddonsAddonSemantic, bool)`

GetSemanticDataOk returns a tuple with the SemanticData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticData

`func (o *AddonsUserAddon) SetSemanticData(v AddonsAddonSemantic)`

SetSemanticData sets SemanticData field to given value.

### HasSemanticData

`func (o *AddonsUserAddon) HasSemanticData() bool`

HasSemanticData returns a boolean if a field has been set.

### GetSensitiveSemantic

`func (o *AddonsUserAddon) GetSensitiveSemantic() map[string]string`

GetSensitiveSemantic returns the SensitiveSemantic field if non-nil, zero value otherwise.

### GetSensitiveSemanticOk

`func (o *AddonsUserAddon) GetSensitiveSemanticOk() (*map[string]string, bool)`

GetSensitiveSemanticOk returns a tuple with the SensitiveSemantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveSemantic

`func (o *AddonsUserAddon) SetSensitiveSemantic(v map[string]string)`

SetSensitiveSemantic sets SensitiveSemantic field to given value.

### HasSensitiveSemantic

`func (o *AddonsUserAddon) HasSensitiveSemantic() bool`

HasSensitiveSemantic returns a boolean if a field has been set.

### GetWidgets

`func (o *AddonsUserAddon) GetWidgets() map[string]interface{}`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *AddonsUserAddon) GetWidgetsOk() (*map[string]interface{}, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *AddonsUserAddon) SetWidgets(v map[string]interface{})`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *AddonsUserAddon) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.

### GetWidgetsSemantic

`func (o *AddonsUserAddon) GetWidgetsSemantic() map[string]interface{}`

GetWidgetsSemantic returns the WidgetsSemantic field if non-nil, zero value otherwise.

### GetWidgetsSemanticOk

`func (o *AddonsUserAddon) GetWidgetsSemanticOk() (*map[string]interface{}, bool)`

GetWidgetsSemanticOk returns a tuple with the WidgetsSemantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetsSemantic

`func (o *AddonsUserAddon) SetWidgetsSemantic(v map[string]interface{})`

SetWidgetsSemantic sets WidgetsSemantic field to given value.

### HasWidgetsSemantic

`func (o *AddonsUserAddon) HasWidgetsSemantic() bool`

HasWidgetsSemantic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


