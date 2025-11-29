# ChatAPIConversationSendMessageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | محتوای پیام متنی برای ارسال | 
**MediaToken** | Pointer to **string** | توکن برای مدیای ضمیمه شده (در صورت وجود) | [optional] 
**ReceiverButtons** | Pointer to [**ChatapiChatButtonGrid**](ChatapiChatButtonGrid.md) |  | [optional] 
**SenderButtons** | Pointer to [**ChatapiChatButtonGrid**](ChatapiChatButtonGrid.md) |  | [optional] 

## Methods

### NewChatAPIConversationSendMessageBody

`func NewChatAPIConversationSendMessageBody(message string, ) *ChatAPIConversationSendMessageBody`

NewChatAPIConversationSendMessageBody instantiates a new ChatAPIConversationSendMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatAPIConversationSendMessageBodyWithDefaults

`func NewChatAPIConversationSendMessageBodyWithDefaults() *ChatAPIConversationSendMessageBody`

NewChatAPIConversationSendMessageBodyWithDefaults instantiates a new ChatAPIConversationSendMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ChatAPIConversationSendMessageBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatAPIConversationSendMessageBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatAPIConversationSendMessageBody) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMediaToken

`func (o *ChatAPIConversationSendMessageBody) GetMediaToken() string`

GetMediaToken returns the MediaToken field if non-nil, zero value otherwise.

### GetMediaTokenOk

`func (o *ChatAPIConversationSendMessageBody) GetMediaTokenOk() (*string, bool)`

GetMediaTokenOk returns a tuple with the MediaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaToken

`func (o *ChatAPIConversationSendMessageBody) SetMediaToken(v string)`

SetMediaToken sets MediaToken field to given value.

### HasMediaToken

`func (o *ChatAPIConversationSendMessageBody) HasMediaToken() bool`

HasMediaToken returns a boolean if a field has been set.

### GetReceiverButtons

`func (o *ChatAPIConversationSendMessageBody) GetReceiverButtons() ChatapiChatButtonGrid`

GetReceiverButtons returns the ReceiverButtons field if non-nil, zero value otherwise.

### GetReceiverButtonsOk

`func (o *ChatAPIConversationSendMessageBody) GetReceiverButtonsOk() (*ChatapiChatButtonGrid, bool)`

GetReceiverButtonsOk returns a tuple with the ReceiverButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverButtons

`func (o *ChatAPIConversationSendMessageBody) SetReceiverButtons(v ChatapiChatButtonGrid)`

SetReceiverButtons sets ReceiverButtons field to given value.

### HasReceiverButtons

`func (o *ChatAPIConversationSendMessageBody) HasReceiverButtons() bool`

HasReceiverButtons returns a boolean if a field has been set.

### GetSenderButtons

`func (o *ChatAPIConversationSendMessageBody) GetSenderButtons() ChatapiChatButtonGrid`

GetSenderButtons returns the SenderButtons field if non-nil, zero value otherwise.

### GetSenderButtonsOk

`func (o *ChatAPIConversationSendMessageBody) GetSenderButtonsOk() (*ChatapiChatButtonGrid, bool)`

GetSenderButtonsOk returns a tuple with the SenderButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderButtons

`func (o *ChatAPIConversationSendMessageBody) SetSenderButtons(v ChatapiChatButtonGrid)`

SetSenderButtons sets SenderButtons field to given value.

### HasSenderButtons

`func (o *ChatAPIConversationSendMessageBody) HasSenderButtons() bool`

HasSenderButtons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


