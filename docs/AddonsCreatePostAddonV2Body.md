# AddonsCreatePostAddonV2Body

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Widgets** | Pointer to [**[]AddonsWidget**](AddonsWidget.md) |  | [optional] 
**Semantic** | Pointer to **map[string]string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 

## Methods

### NewAddonsCreatePostAddonV2Body

`func NewAddonsCreatePostAddonV2Body() *AddonsCreatePostAddonV2Body`

NewAddonsCreatePostAddonV2Body instantiates a new AddonsCreatePostAddonV2Body object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsCreatePostAddonV2BodyWithDefaults

`func NewAddonsCreatePostAddonV2BodyWithDefaults() *AddonsCreatePostAddonV2Body`

NewAddonsCreatePostAddonV2BodyWithDefaults instantiates a new AddonsCreatePostAddonV2Body object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidgets

`func (o *AddonsCreatePostAddonV2Body) GetWidgets() []AddonsWidget`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *AddonsCreatePostAddonV2Body) GetWidgetsOk() (*[]AddonsWidget, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *AddonsCreatePostAddonV2Body) SetWidgets(v []AddonsWidget)`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *AddonsCreatePostAddonV2Body) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.

### GetSemantic

`func (o *AddonsCreatePostAddonV2Body) GetSemantic() map[string]string`

GetSemantic returns the Semantic field if non-nil, zero value otherwise.

### GetSemanticOk

`func (o *AddonsCreatePostAddonV2Body) GetSemanticOk() (*map[string]string, bool)`

GetSemanticOk returns a tuple with the Semantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemantic

`func (o *AddonsCreatePostAddonV2Body) SetSemantic(v map[string]string)`

SetSemantic sets Semantic field to given value.

### HasSemantic

`func (o *AddonsCreatePostAddonV2Body) HasSemantic() bool`

HasSemantic returns a boolean if a field has been set.

### GetNotes

`func (o *AddonsCreatePostAddonV2Body) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AddonsCreatePostAddonV2Body) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AddonsCreatePostAddonV2Body) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AddonsCreatePostAddonV2Body) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


