# EventsEventButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**AddonsAction**](AddonsAction.md) |  | [optional] 
**Title** | Pointer to **string** | متن برای نمایش روی دکمه | [optional] 

## Methods

### NewEventsEventButton

`func NewEventsEventButton() *EventsEventButton`

NewEventsEventButton instantiates a new EventsEventButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsEventButtonWithDefaults

`func NewEventsEventButtonWithDefaults() *EventsEventButton`

NewEventsEventButtonWithDefaults instantiates a new EventsEventButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *EventsEventButton) GetAction() AddonsAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EventsEventButton) GetActionOk() (*AddonsAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EventsEventButton) SetAction(v AddonsAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *EventsEventButton) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetTitle

`func (o *EventsEventButton) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventsEventButton) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventsEventButton) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EventsEventButton) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


