# EventsRegisterEventSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventResourceId** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to [**EventsRegisterEventSubscriptionRequestEventType**](EventsRegisterEventSubscriptionRequestEventType.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEventsRegisterEventSubscriptionRequest

`func NewEventsRegisterEventSubscriptionRequest() *EventsRegisterEventSubscriptionRequest`

NewEventsRegisterEventSubscriptionRequest instantiates a new EventsRegisterEventSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsRegisterEventSubscriptionRequestWithDefaults

`func NewEventsRegisterEventSubscriptionRequestWithDefaults() *EventsRegisterEventSubscriptionRequest`

NewEventsRegisterEventSubscriptionRequestWithDefaults instantiates a new EventsRegisterEventSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventResourceId

`func (o *EventsRegisterEventSubscriptionRequest) GetEventResourceId() string`

GetEventResourceId returns the EventResourceId field if non-nil, zero value otherwise.

### GetEventResourceIdOk

`func (o *EventsRegisterEventSubscriptionRequest) GetEventResourceIdOk() (*string, bool)`

GetEventResourceIdOk returns a tuple with the EventResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventResourceId

`func (o *EventsRegisterEventSubscriptionRequest) SetEventResourceId(v string)`

SetEventResourceId sets EventResourceId field to given value.

### HasEventResourceId

`func (o *EventsRegisterEventSubscriptionRequest) HasEventResourceId() bool`

HasEventResourceId returns a boolean if a field has been set.

### GetEventType

`func (o *EventsRegisterEventSubscriptionRequest) GetEventType() EventsRegisterEventSubscriptionRequestEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventsRegisterEventSubscriptionRequest) GetEventTypeOk() (*EventsRegisterEventSubscriptionRequestEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventsRegisterEventSubscriptionRequest) SetEventType(v EventsRegisterEventSubscriptionRequestEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventsRegisterEventSubscriptionRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetMetadata

`func (o *EventsRegisterEventSubscriptionRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EventsRegisterEventSubscriptionRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EventsRegisterEventSubscriptionRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EventsRegisterEventSubscriptionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


