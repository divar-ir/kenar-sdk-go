# AddonsCreateUserAddonV2Body

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Widgets** | Pointer to [**[]AddonsWidget**](AddonsWidget.md) |  | [optional] 
**Semantic** | Pointer to **map[string]string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**TicketUuid** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **int32** |  | [optional] 

## Methods

### NewAddonsCreateUserAddonV2Body

`func NewAddonsCreateUserAddonV2Body() *AddonsCreateUserAddonV2Body`

NewAddonsCreateUserAddonV2Body instantiates a new AddonsCreateUserAddonV2Body object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsCreateUserAddonV2BodyWithDefaults

`func NewAddonsCreateUserAddonV2BodyWithDefaults() *AddonsCreateUserAddonV2Body`

NewAddonsCreateUserAddonV2BodyWithDefaults instantiates a new AddonsCreateUserAddonV2Body object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidgets

`func (o *AddonsCreateUserAddonV2Body) GetWidgets() []AddonsWidget`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *AddonsCreateUserAddonV2Body) GetWidgetsOk() (*[]AddonsWidget, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *AddonsCreateUserAddonV2Body) SetWidgets(v []AddonsWidget)`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *AddonsCreateUserAddonV2Body) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.

### GetSemantic

`func (o *AddonsCreateUserAddonV2Body) GetSemantic() map[string]string`

GetSemantic returns the Semantic field if non-nil, zero value otherwise.

### GetSemanticOk

`func (o *AddonsCreateUserAddonV2Body) GetSemanticOk() (*map[string]string, bool)`

GetSemanticOk returns a tuple with the Semantic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemantic

`func (o *AddonsCreateUserAddonV2Body) SetSemantic(v map[string]string)`

SetSemantic sets Semantic field to given value.

### HasSemantic

`func (o *AddonsCreateUserAddonV2Body) HasSemantic() bool`

HasSemantic returns a boolean if a field has been set.

### GetPhone

`func (o *AddonsCreateUserAddonV2Body) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddonsCreateUserAddonV2Body) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddonsCreateUserAddonV2Body) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddonsCreateUserAddonV2Body) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCategories

`func (o *AddonsCreateUserAddonV2Body) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AddonsCreateUserAddonV2Body) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AddonsCreateUserAddonV2Body) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AddonsCreateUserAddonV2Body) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetTicketUuid

`func (o *AddonsCreateUserAddonV2Body) GetTicketUuid() string`

GetTicketUuid returns the TicketUuid field if non-nil, zero value otherwise.

### GetTicketUuidOk

`func (o *AddonsCreateUserAddonV2Body) GetTicketUuidOk() (*string, bool)`

GetTicketUuidOk returns a tuple with the TicketUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketUuid

`func (o *AddonsCreateUserAddonV2Body) SetTicketUuid(v string)`

SetTicketUuid sets TicketUuid field to given value.

### HasTicketUuid

`func (o *AddonsCreateUserAddonV2Body) HasTicketUuid() bool`

HasTicketUuid returns a boolean if a field has been set.

### GetCost

`func (o *AddonsCreateUserAddonV2Body) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AddonsCreateUserAddonV2Body) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AddonsCreateUserAddonV2Body) SetCost(v int32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AddonsCreateUserAddonV2Body) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


