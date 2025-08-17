# WidgetsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PagePopLink** | Pointer to **bool** | deprecated; use is_new_flow &#x3D; 1 in PagePresentation instead. | [optional] 
**Payload** | Pointer to [**ProtobufAny**](ProtobufAny.md) |  | [optional] 
**Type** | Pointer to [**WidgetsActionType**](WidgetsActionType.md) |  | [optional] 

## Methods

### NewWidgetsAction

`func NewWidgetsAction() *WidgetsAction`

NewWidgetsAction instantiates a new WidgetsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetsActionWithDefaults

`func NewWidgetsActionWithDefaults() *WidgetsAction`

NewWidgetsActionWithDefaults instantiates a new WidgetsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagePopLink

`func (o *WidgetsAction) GetPagePopLink() bool`

GetPagePopLink returns the PagePopLink field if non-nil, zero value otherwise.

### GetPagePopLinkOk

`func (o *WidgetsAction) GetPagePopLinkOk() (*bool, bool)`

GetPagePopLinkOk returns a tuple with the PagePopLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagePopLink

`func (o *WidgetsAction) SetPagePopLink(v bool)`

SetPagePopLink sets PagePopLink field to given value.

### HasPagePopLink

`func (o *WidgetsAction) HasPagePopLink() bool`

HasPagePopLink returns a boolean if a field has been set.

### GetPayload

`func (o *WidgetsAction) GetPayload() ProtobufAny`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WidgetsAction) GetPayloadOk() (*ProtobufAny, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WidgetsAction) SetPayload(v ProtobufAny)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *WidgetsAction) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetType

`func (o *WidgetsAction) GetType() WidgetsActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WidgetsAction) GetTypeOk() (*WidgetsActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WidgetsAction) SetType(v WidgetsActionType)`

SetType sets Type field to given value.

### HasType

`func (o *WidgetsAction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


