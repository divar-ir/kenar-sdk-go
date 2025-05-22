# AddonsCreatePostAddonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkInSpec** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Selector** | Pointer to [**AddonsAddonSelector**](AddonsAddonSelector.md) |  | [optional] 
**Semantic** | Pointer to **map[string]string** |  | [optional] 
**SemanticSensitives** | Pointer to **[]string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Widgets** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAddonsCreatePostAddonRequest

`func NewAddonsCreatePostAddonRequest() *AddonsCreatePostAddonRequest`

NewAddonsCreatePostAddonRequest instantiates a new AddonsCreatePostAddonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsCreatePostAddonRequestWithDefaults

`func NewAddonsCreatePostAddonRequestWithDefaults() *AddonsCreatePostAddonRequest`

NewAddonsCreatePostAddonRequestWithDefaults instantiates a new AddonsCreatePostAddonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkInSpec

`func (o *AddonsCreatePostAddonRequest) GetLinkInSpec() string`

GetLinkInSpec returns the LinkInSpec field if non-nil, zero value otherwise.

### GetLinkInSpecOk

`func (o *AddonsCreatePostAddonRequest) GetLinkInSpecOk() (*string, bool)`

GetLinkInSpecOk returns a tuple with the LinkInSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkInSpec

`func (o *AddonsCreatePostAddonRequest) SetLinkInSpec(v string)`

SetLinkInSpec sets LinkInSpec field to given value.

### HasLinkInSpec

`func (o *AddonsCreatePostAddonRequest) HasLinkInSpec() bool`

HasLinkInSpec returns a boolean if a field has been set.

### GetNotes

`func (o *AddonsCreatePostAddonRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AddonsCreatePostAddonRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AddonsCreatePostAddonRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AddonsCreatePostAddonRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetSelector

`func (o *AddonsCreatePostAddonRequest) GetSelector() AddonsAddonSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *AddonsCreatePostAddonRequest) GetSelectorOk() (*AddonsAddonSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *AddonsCreatePostAddonRequest) SetSelector(v AddonsAddonSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *AddonsCreatePostAddonRequest) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSemantic

`func (o *AddonsCreatePostAddonRequest) GetSemantic() map[string]string`

GetSemantic returns the Semantic field if non-nil, zero value otherwise.

### GetSemanticOk

`func (o *AddonsCreatePostAddonRequest) GetSemanticOk() (*map[string]string, bool)`

GetSemanticOk returns a tuple with the Semantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemantic

`func (o *AddonsCreatePostAddonRequest) SetSemantic(v map[string]string)`

SetSemantic sets Semantic field to given value.

### HasSemantic

`func (o *AddonsCreatePostAddonRequest) HasSemantic() bool`

HasSemantic returns a boolean if a field has been set.

### GetSemanticSensitives

`func (o *AddonsCreatePostAddonRequest) GetSemanticSensitives() []string`

GetSemanticSensitives returns the SemanticSensitives field if non-nil, zero value otherwise.

### GetSemanticSensitivesOk

`func (o *AddonsCreatePostAddonRequest) GetSemanticSensitivesOk() (*[]string, bool)`

GetSemanticSensitivesOk returns a tuple with the SemanticSensitives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticSensitives

`func (o *AddonsCreatePostAddonRequest) SetSemanticSensitives(v []string)`

SetSemanticSensitives sets SemanticSensitives field to given value.

### HasSemanticSensitives

`func (o *AddonsCreatePostAddonRequest) HasSemanticSensitives() bool`

HasSemanticSensitives returns a boolean if a field has been set.

### GetToken

`func (o *AddonsCreatePostAddonRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddonsCreatePostAddonRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddonsCreatePostAddonRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AddonsCreatePostAddonRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetWidgets

`func (o *AddonsCreatePostAddonRequest) GetWidgets() map[string]interface{}`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *AddonsCreatePostAddonRequest) GetWidgetsOk() (*map[string]interface{}, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *AddonsCreatePostAddonRequest) SetWidgets(v map[string]interface{})`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *AddonsCreatePostAddonRequest) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


