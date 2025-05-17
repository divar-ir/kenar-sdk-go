# AddonsPostAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaData** | Pointer to [**AddonsAddonMetaData**](AddonsAddonMetaData.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**App** | Pointer to [**AppsApp**](AppsApp.md) |  | [optional] 
**Widgets** | Pointer to **map[string]interface{}** |  | [optional] 
**Score** | Pointer to **string** |  | [optional] 
**Selector** | Pointer to [**AddonsAddonSelector**](AddonsAddonSelector.md) |  | [optional] 
**Linkage** | Pointer to [**AddonsAddonLinkage**](AddonsAddonLinkage.md) |  | [optional] 
**SecondaryLinks** | Pointer to [**AddonsAddonSecondaryLinks**](AddonsAddonSecondaryLinks.md) |  | [optional] 
**Semantic** | Pointer to **map[string]string** |  | [optional] 
**SemanticData** | Pointer to [**AddonsAddonSemantic**](AddonsAddonSemantic.md) |  | [optional] 
**SensitiveSemantic** | Pointer to **map[string]string** |  | [optional] 
**WidgetsSemantic** | Pointer to **map[string]interface{}** |  | [optional] 
**SemanticSensitives** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddonsPostAddon

`func NewAddonsPostAddon() *AddonsPostAddon`

NewAddonsPostAddon instantiates a new AddonsPostAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsPostAddonWithDefaults

`func NewAddonsPostAddonWithDefaults() *AddonsPostAddon`

NewAddonsPostAddonWithDefaults instantiates a new AddonsPostAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaData

`func (o *AddonsPostAddon) GetMetaData() AddonsAddonMetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AddonsPostAddon) GetMetaDataOk() (*AddonsAddonMetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AddonsPostAddon) SetMetaData(v AddonsAddonMetaData)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AddonsPostAddon) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetToken

`func (o *AddonsPostAddon) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddonsPostAddon) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddonsPostAddon) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AddonsPostAddon) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetApp

`func (o *AddonsPostAddon) GetApp() AppsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AddonsPostAddon) GetAppOk() (*AppsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AddonsPostAddon) SetApp(v AppsApp)`

SetApp sets App field to given value.

### HasApp

`func (o *AddonsPostAddon) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetWidgets

`func (o *AddonsPostAddon) GetWidgets() map[string]interface{}`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *AddonsPostAddon) GetWidgetsOk() (*map[string]interface{}, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *AddonsPostAddon) SetWidgets(v map[string]interface{})`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *AddonsPostAddon) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.

### GetScore

`func (o *AddonsPostAddon) GetScore() string`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AddonsPostAddon) GetScoreOk() (*string, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AddonsPostAddon) SetScore(v string)`

SetScore sets Score field to given value.

### HasScore

`func (o *AddonsPostAddon) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSelector

`func (o *AddonsPostAddon) GetSelector() AddonsAddonSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *AddonsPostAddon) GetSelectorOk() (*AddonsAddonSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *AddonsPostAddon) SetSelector(v AddonsAddonSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *AddonsPostAddon) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetLinkage

`func (o *AddonsPostAddon) GetLinkage() AddonsAddonLinkage`

GetLinkage returns the Linkage field if non-nil, zero value otherwise.

### GetLinkageOk

`func (o *AddonsPostAddon) GetLinkageOk() (*AddonsAddonLinkage, bool)`

GetLinkageOk returns a tuple with the Linkage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkage

`func (o *AddonsPostAddon) SetLinkage(v AddonsAddonLinkage)`

SetLinkage sets Linkage field to given value.

### HasLinkage

`func (o *AddonsPostAddon) HasLinkage() bool`

HasLinkage returns a boolean if a field has been set.

### GetSecondaryLinks

`func (o *AddonsPostAddon) GetSecondaryLinks() AddonsAddonSecondaryLinks`

GetSecondaryLinks returns the SecondaryLinks field if non-nil, zero value otherwise.

### GetSecondaryLinksOk

`func (o *AddonsPostAddon) GetSecondaryLinksOk() (*AddonsAddonSecondaryLinks, bool)`

GetSecondaryLinksOk returns a tuple with the SecondaryLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryLinks

`func (o *AddonsPostAddon) SetSecondaryLinks(v AddonsAddonSecondaryLinks)`

SetSecondaryLinks sets SecondaryLinks field to given value.

### HasSecondaryLinks

`func (o *AddonsPostAddon) HasSecondaryLinks() bool`

HasSecondaryLinks returns a boolean if a field has been set.

### GetSemantic

`func (o *AddonsPostAddon) GetSemantic() map[string]string`

GetSemantic returns the Semantic field if non-nil, zero value otherwise.

### GetSemanticOk

`func (o *AddonsPostAddon) GetSemanticOk() (*map[string]string, bool)`

GetSemanticOk returns a tuple with the Semantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemantic

`func (o *AddonsPostAddon) SetSemantic(v map[string]string)`

SetSemantic sets Semantic field to given value.

### HasSemantic

`func (o *AddonsPostAddon) HasSemantic() bool`

HasSemantic returns a boolean if a field has been set.

### GetSemanticData

`func (o *AddonsPostAddon) GetSemanticData() AddonsAddonSemantic`

GetSemanticData returns the SemanticData field if non-nil, zero value otherwise.

### GetSemanticDataOk

`func (o *AddonsPostAddon) GetSemanticDataOk() (*AddonsAddonSemantic, bool)`

GetSemanticDataOk returns a tuple with the SemanticData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticData

`func (o *AddonsPostAddon) SetSemanticData(v AddonsAddonSemantic)`

SetSemanticData sets SemanticData field to given value.

### HasSemanticData

`func (o *AddonsPostAddon) HasSemanticData() bool`

HasSemanticData returns a boolean if a field has been set.

### GetSensitiveSemantic

`func (o *AddonsPostAddon) GetSensitiveSemantic() map[string]string`

GetSensitiveSemantic returns the SensitiveSemantic field if non-nil, zero value otherwise.

### GetSensitiveSemanticOk

`func (o *AddonsPostAddon) GetSensitiveSemanticOk() (*map[string]string, bool)`

GetSensitiveSemanticOk returns a tuple with the SensitiveSemantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveSemantic

`func (o *AddonsPostAddon) SetSensitiveSemantic(v map[string]string)`

SetSensitiveSemantic sets SensitiveSemantic field to given value.

### HasSensitiveSemantic

`func (o *AddonsPostAddon) HasSensitiveSemantic() bool`

HasSensitiveSemantic returns a boolean if a field has been set.

### GetWidgetsSemantic

`func (o *AddonsPostAddon) GetWidgetsSemantic() map[string]interface{}`

GetWidgetsSemantic returns the WidgetsSemantic field if non-nil, zero value otherwise.

### GetWidgetsSemanticOk

`func (o *AddonsPostAddon) GetWidgetsSemanticOk() (*map[string]interface{}, bool)`

GetWidgetsSemanticOk returns a tuple with the WidgetsSemantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetsSemantic

`func (o *AddonsPostAddon) SetWidgetsSemantic(v map[string]interface{})`

SetWidgetsSemantic sets WidgetsSemantic field to given value.

### HasWidgetsSemantic

`func (o *AddonsPostAddon) HasWidgetsSemantic() bool`

HasWidgetsSemantic returns a boolean if a field has been set.

### GetSemanticSensitives

`func (o *AddonsPostAddon) GetSemanticSensitives() []string`

GetSemanticSensitives returns the SemanticSensitives field if non-nil, zero value otherwise.

### GetSemanticSensitivesOk

`func (o *AddonsPostAddon) GetSemanticSensitivesOk() (*[]string, bool)`

GetSemanticSensitivesOk returns a tuple with the SemanticSensitives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticSensitives

`func (o *AddonsPostAddon) SetSemanticSensitives(v []string)`

SetSemanticSensitives sets SemanticSensitives field to given value.

### HasSemanticSensitives

`func (o *AddonsPostAddon) HasSemanticSensitives() bool`

HasSemanticSensitives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


