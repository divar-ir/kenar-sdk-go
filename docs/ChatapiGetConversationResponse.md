# ChatapiGetConversationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversation** | [**ChatapiConversation**](ChatapiConversation.md) |  | 
**Messages** | [**[]ChatapiMessage**](ChatapiMessage.md) | فهرست پیام‌های مکالمه | 

## Methods

### NewChatapiGetConversationResponse

`func NewChatapiGetConversationResponse(conversation ChatapiConversation, messages []ChatapiMessage, ) *ChatapiGetConversationResponse`

NewChatapiGetConversationResponse instantiates a new ChatapiGetConversationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatapiGetConversationResponseWithDefaults

`func NewChatapiGetConversationResponseWithDefaults() *ChatapiGetConversationResponse`

NewChatapiGetConversationResponseWithDefaults instantiates a new ChatapiGetConversationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversation

`func (o *ChatapiGetConversationResponse) GetConversation() ChatapiConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *ChatapiGetConversationResponse) GetConversationOk() (*ChatapiConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *ChatapiGetConversationResponse) SetConversation(v ChatapiConversation)`

SetConversation sets Conversation field to given value.


### GetMessages

`func (o *ChatapiGetConversationResponse) GetMessages() []ChatapiMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatapiGetConversationResponse) GetMessagesOk() (*[]ChatapiMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatapiGetConversationResponse) SetMessages(v []ChatapiMessage)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


