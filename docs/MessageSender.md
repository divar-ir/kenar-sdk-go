# MessageSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Side** | Pointer to [**ChatapiMessageSenderSide**](ChatapiMessageSenderSide.md) |  | [optional] 
**Type** | Pointer to [**ChatapiMessageSenderType**](ChatapiMessageSenderType.md) |  | [optional] 

## Methods

### NewMessageSender

`func NewMessageSender() *MessageSender`

NewMessageSender instantiates a new MessageSender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageSenderWithDefaults

`func NewMessageSenderWithDefaults() *MessageSender`

NewMessageSenderWithDefaults instantiates a new MessageSender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSide

`func (o *MessageSender) GetSide() ChatapiMessageSenderSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *MessageSender) GetSideOk() (*ChatapiMessageSenderSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *MessageSender) SetSide(v ChatapiMessageSenderSide)`

SetSide sets Side field to given value.

### HasSide

`func (o *MessageSender) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetType

`func (o *MessageSender) GetType() ChatapiMessageSenderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageSender) GetTypeOk() (*ChatapiMessageSenderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageSender) SetType(v ChatapiMessageSenderType)`

SetType sets Type field to given value.

### HasType

`func (o *MessageSender) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


