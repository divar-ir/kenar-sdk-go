# AddonsWidget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventRow** | Pointer to [**AddonsEventRow**](AddonsEventRow.md) |  | [optional] 
**ButtonBar** | Pointer to [**AddonsButtonBar**](AddonsButtonBar.md) |  | [optional] 
**TitleRow** | Pointer to [**AddonsTitleRow**](AddonsTitleRow.md) |  | [optional] 
**SubtitleRow** | Pointer to [**AddonsSubtitleRow**](AddonsSubtitleRow.md) |  | [optional] 
**SelectorRow** | Pointer to [**AddonsSelectorRow**](AddonsSelectorRow.md) |  | [optional] 
**ScoreRow** | Pointer to [**AddonsScoreRow**](AddonsScoreRow.md) |  | [optional] 
**ImageCarouselRow** | Pointer to [**AddonsImageCarouselRow**](AddonsImageCarouselRow.md) |  | [optional] 
**GroupInfoRow** | Pointer to [**AddonsGroupInfoRow**](AddonsGroupInfoRow.md) |  | [optional] 
**EvaluationRow** | Pointer to [**AddonsEvaluationRow**](AddonsEvaluationRow.md) |  | [optional] 
**DescriptionRow** | Pointer to [**AddonsDescriptionRow**](AddonsDescriptionRow.md) |  | [optional] 
**SemanticPaths** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAddonsWidget

`func NewAddonsWidget() *AddonsWidget`

NewAddonsWidget instantiates a new AddonsWidget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonsWidgetWithDefaults

`func NewAddonsWidgetWithDefaults() *AddonsWidget`

NewAddonsWidgetWithDefaults instantiates a new AddonsWidget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventRow

`func (o *AddonsWidget) GetEventRow() AddonsEventRow`

GetEventRow returns the EventRow field if non-nil, zero value otherwise.

### GetEventRowOk

`func (o *AddonsWidget) GetEventRowOk() (*AddonsEventRow, bool)`

GetEventRowOk returns a tuple with the EventRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRow

`func (o *AddonsWidget) SetEventRow(v AddonsEventRow)`

SetEventRow sets EventRow field to given value.

### HasEventRow

`func (o *AddonsWidget) HasEventRow() bool`

HasEventRow returns a boolean if a field has been set.

### GetButtonBar

`func (o *AddonsWidget) GetButtonBar() AddonsButtonBar`

GetButtonBar returns the ButtonBar field if non-nil, zero value otherwise.

### GetButtonBarOk

`func (o *AddonsWidget) GetButtonBarOk() (*AddonsButtonBar, bool)`

GetButtonBarOk returns a tuple with the ButtonBar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonBar

`func (o *AddonsWidget) SetButtonBar(v AddonsButtonBar)`

SetButtonBar sets ButtonBar field to given value.

### HasButtonBar

`func (o *AddonsWidget) HasButtonBar() bool`

HasButtonBar returns a boolean if a field has been set.

### GetTitleRow

`func (o *AddonsWidget) GetTitleRow() AddonsTitleRow`

GetTitleRow returns the TitleRow field if non-nil, zero value otherwise.

### GetTitleRowOk

`func (o *AddonsWidget) GetTitleRowOk() (*AddonsTitleRow, bool)`

GetTitleRowOk returns a tuple with the TitleRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleRow

`func (o *AddonsWidget) SetTitleRow(v AddonsTitleRow)`

SetTitleRow sets TitleRow field to given value.

### HasTitleRow

`func (o *AddonsWidget) HasTitleRow() bool`

HasTitleRow returns a boolean if a field has been set.

### GetSubtitleRow

`func (o *AddonsWidget) GetSubtitleRow() AddonsSubtitleRow`

GetSubtitleRow returns the SubtitleRow field if non-nil, zero value otherwise.

### GetSubtitleRowOk

`func (o *AddonsWidget) GetSubtitleRowOk() (*AddonsSubtitleRow, bool)`

GetSubtitleRowOk returns a tuple with the SubtitleRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleRow

`func (o *AddonsWidget) SetSubtitleRow(v AddonsSubtitleRow)`

SetSubtitleRow sets SubtitleRow field to given value.

### HasSubtitleRow

`func (o *AddonsWidget) HasSubtitleRow() bool`

HasSubtitleRow returns a boolean if a field has been set.

### GetSelectorRow

`func (o *AddonsWidget) GetSelectorRow() AddonsSelectorRow`

GetSelectorRow returns the SelectorRow field if non-nil, zero value otherwise.

### GetSelectorRowOk

`func (o *AddonsWidget) GetSelectorRowOk() (*AddonsSelectorRow, bool)`

GetSelectorRowOk returns a tuple with the SelectorRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorRow

`func (o *AddonsWidget) SetSelectorRow(v AddonsSelectorRow)`

SetSelectorRow sets SelectorRow field to given value.

### HasSelectorRow

`func (o *AddonsWidget) HasSelectorRow() bool`

HasSelectorRow returns a boolean if a field has been set.

### GetScoreRow

`func (o *AddonsWidget) GetScoreRow() AddonsScoreRow`

GetScoreRow returns the ScoreRow field if non-nil, zero value otherwise.

### GetScoreRowOk

`func (o *AddonsWidget) GetScoreRowOk() (*AddonsScoreRow, bool)`

GetScoreRowOk returns a tuple with the ScoreRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreRow

`func (o *AddonsWidget) SetScoreRow(v AddonsScoreRow)`

SetScoreRow sets ScoreRow field to given value.

### HasScoreRow

`func (o *AddonsWidget) HasScoreRow() bool`

HasScoreRow returns a boolean if a field has been set.

### GetImageCarouselRow

`func (o *AddonsWidget) GetImageCarouselRow() AddonsImageCarouselRow`

GetImageCarouselRow returns the ImageCarouselRow field if non-nil, zero value otherwise.

### GetImageCarouselRowOk

`func (o *AddonsWidget) GetImageCarouselRowOk() (*AddonsImageCarouselRow, bool)`

GetImageCarouselRowOk returns a tuple with the ImageCarouselRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCarouselRow

`func (o *AddonsWidget) SetImageCarouselRow(v AddonsImageCarouselRow)`

SetImageCarouselRow sets ImageCarouselRow field to given value.

### HasImageCarouselRow

`func (o *AddonsWidget) HasImageCarouselRow() bool`

HasImageCarouselRow returns a boolean if a field has been set.

### GetGroupInfoRow

`func (o *AddonsWidget) GetGroupInfoRow() AddonsGroupInfoRow`

GetGroupInfoRow returns the GroupInfoRow field if non-nil, zero value otherwise.

### GetGroupInfoRowOk

`func (o *AddonsWidget) GetGroupInfoRowOk() (*AddonsGroupInfoRow, bool)`

GetGroupInfoRowOk returns a tuple with the GroupInfoRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupInfoRow

`func (o *AddonsWidget) SetGroupInfoRow(v AddonsGroupInfoRow)`

SetGroupInfoRow sets GroupInfoRow field to given value.

### HasGroupInfoRow

`func (o *AddonsWidget) HasGroupInfoRow() bool`

HasGroupInfoRow returns a boolean if a field has been set.

### GetEvaluationRow

`func (o *AddonsWidget) GetEvaluationRow() AddonsEvaluationRow`

GetEvaluationRow returns the EvaluationRow field if non-nil, zero value otherwise.

### GetEvaluationRowOk

`func (o *AddonsWidget) GetEvaluationRowOk() (*AddonsEvaluationRow, bool)`

GetEvaluationRowOk returns a tuple with the EvaluationRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationRow

`func (o *AddonsWidget) SetEvaluationRow(v AddonsEvaluationRow)`

SetEvaluationRow sets EvaluationRow field to given value.

### HasEvaluationRow

`func (o *AddonsWidget) HasEvaluationRow() bool`

HasEvaluationRow returns a boolean if a field has been set.

### GetDescriptionRow

`func (o *AddonsWidget) GetDescriptionRow() AddonsDescriptionRow`

GetDescriptionRow returns the DescriptionRow field if non-nil, zero value otherwise.

### GetDescriptionRowOk

`func (o *AddonsWidget) GetDescriptionRowOk() (*AddonsDescriptionRow, bool)`

GetDescriptionRowOk returns a tuple with the DescriptionRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionRow

`func (o *AddonsWidget) SetDescriptionRow(v AddonsDescriptionRow)`

SetDescriptionRow sets DescriptionRow field to given value.

### HasDescriptionRow

`func (o *AddonsWidget) HasDescriptionRow() bool`

HasDescriptionRow returns a boolean if a field has been set.

### GetSemanticPaths

`func (o *AddonsWidget) GetSemanticPaths() map[string]string`

GetSemanticPaths returns the SemanticPaths field if non-nil, zero value otherwise.

### GetSemanticPathsOk

`func (o *AddonsWidget) GetSemanticPathsOk() (*map[string]string, bool)`

GetSemanticPathsOk returns a tuple with the SemanticPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticPaths

`func (o *AddonsWidget) SetSemanticPaths(v map[string]string)`

SetSemanticPaths sets SemanticPaths field to given value.

### HasSemanticPaths

`func (o *AddonsWidget) HasSemanticPaths() bool`

HasSemanticPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


