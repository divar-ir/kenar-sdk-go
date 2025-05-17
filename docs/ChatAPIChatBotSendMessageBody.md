# ChatAPIChatBotSendMessageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**TextMessage** | Pointer to **string** |  | [optional] 
**MediaToken** | Pointer to **string** |  | [optional] 
**Buttons** | Pointer to [**ChatapiChatButtonGrid**](ChatapiChatButtonGrid.md) |  | [optional] 

## Methods

### NewChatAPIChatBotSendMessageBody

`func NewChatAPIChatBotSendMessageBody() *ChatAPIChatBotSendMessageBody`

NewChatAPIChatBotSendMessageBody instantiates a new ChatAPIChatBotSendMessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatAPIChatBotSendMessageBodyWithDefaults

`func NewChatAPIChatBotSendMessageBodyWithDefaults() *ChatAPIChatBotSendMessageBody`

NewChatAPIChatBotSendMessageBodyWithDefaults instantiates a new ChatAPIChatBotSendMessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ChatAPIChatBotSendMessageBody) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ChatAPIChatBotSendMessageBody) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ChatAPIChatBotSendMessageBody) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ChatAPIChatBotSendMessageBody) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTextMessage

`func (o *ChatAPIChatBotSendMessageBody) GetTextMessage() string`

GetTextMessage returns the TextMessage field if non-nil, zero value otherwise.

### GetTextMessageOk

`func (o *ChatAPIChatBotSendMessageBody) GetTextMessageOk() (*string, bool)`

GetTextMessageOk returns a tuple with the TextMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextMessage

`func (o *ChatAPIChatBotSendMessageBody) SetTextMessage(v string)`

SetTextMessage sets TextMessage field to given value.

### HasTextMessage

`func (o *ChatAPIChatBotSendMessageBody) HasTextMessage() bool`

HasTextMessage returns a boolean if a field has been set.

### GetMediaToken

`func (o *ChatAPIChatBotSendMessageBody) GetMediaToken() string`

GetMediaToken returns the MediaToken field if non-nil, zero value otherwise.

### GetMediaTokenOk

`func (o *ChatAPIChatBotSendMessageBody) GetMediaTokenOk() (*string, bool)`

GetMediaTokenOk returns a tuple with the MediaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaToken

`func (o *ChatAPIChatBotSendMessageBody) SetMediaToken(v string)`

SetMediaToken sets MediaToken field to given value.

### HasMediaToken

`func (o *ChatAPIChatBotSendMessageBody) HasMediaToken() bool`

HasMediaToken returns a boolean if a field has been set.

### GetButtons

`func (o *ChatAPIChatBotSendMessageBody) GetButtons() ChatapiChatButtonGrid`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *ChatAPIChatBotSendMessageBody) GetButtonsOk() (*ChatapiChatButtonGrid, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *ChatAPIChatBotSendMessageBody) SetButtons(v ChatapiChatButtonGrid)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *ChatAPIChatBotSendMessageBody) HasButtons() bool`

HasButtons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


