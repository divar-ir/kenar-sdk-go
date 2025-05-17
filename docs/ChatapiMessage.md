# ChatapiMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Conversation** | Pointer to [**ChatapiConversation**](ChatapiConversation.md) |  | [optional] 
**Sender** | Pointer to [**MessageSender**](MessageSender.md) |  | [optional] 
**Type** | Pointer to [**ChatapiMessageType**](ChatapiMessageType.md) |  | [optional] 
**SentAt** | Pointer to **time.Time** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**ImageData** | Pointer to [**MessageImageData**](MessageImageData.md) |  | [optional] 
**FileData** | Pointer to [**MessageFileData**](MessageFileData.md) |  | [optional] 
**VoiceData** | Pointer to [**MessageVoiceData**](MessageVoiceData.md) |  | [optional] 
**LocationData** | Pointer to [**MessageLocationData**](MessageLocationData.md) |  | [optional] 
**VideoData** | Pointer to [**MessageVideoData**](MessageVideoData.md) |  | [optional] 

## Methods

### NewChatapiMessage

`func NewChatapiMessage() *ChatapiMessage`

NewChatapiMessage instantiates a new ChatapiMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatapiMessageWithDefaults

`func NewChatapiMessageWithDefaults() *ChatapiMessage`

NewChatapiMessageWithDefaults instantiates a new ChatapiMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatapiMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatapiMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatapiMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatapiMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConversation

`func (o *ChatapiMessage) GetConversation() ChatapiConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *ChatapiMessage) GetConversationOk() (*ChatapiConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *ChatapiMessage) SetConversation(v ChatapiConversation)`

SetConversation sets Conversation field to given value.

### HasConversation

`func (o *ChatapiMessage) HasConversation() bool`

HasConversation returns a boolean if a field has been set.

### GetSender

`func (o *ChatapiMessage) GetSender() MessageSender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ChatapiMessage) GetSenderOk() (*MessageSender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ChatapiMessage) SetSender(v MessageSender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *ChatapiMessage) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetType

`func (o *ChatapiMessage) GetType() ChatapiMessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChatapiMessage) GetTypeOk() (*ChatapiMessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChatapiMessage) SetType(v ChatapiMessageType)`

SetType sets Type field to given value.

### HasType

`func (o *ChatapiMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSentAt

`func (o *ChatapiMessage) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *ChatapiMessage) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *ChatapiMessage) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *ChatapiMessage) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetText

`func (o *ChatapiMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChatapiMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChatapiMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ChatapiMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetImageData

`func (o *ChatapiMessage) GetImageData() MessageImageData`

GetImageData returns the ImageData field if non-nil, zero value otherwise.

### GetImageDataOk

`func (o *ChatapiMessage) GetImageDataOk() (*MessageImageData, bool)`

GetImageDataOk returns a tuple with the ImageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageData

`func (o *ChatapiMessage) SetImageData(v MessageImageData)`

SetImageData sets ImageData field to given value.

### HasImageData

`func (o *ChatapiMessage) HasImageData() bool`

HasImageData returns a boolean if a field has been set.

### GetFileData

`func (o *ChatapiMessage) GetFileData() MessageFileData`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *ChatapiMessage) GetFileDataOk() (*MessageFileData, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *ChatapiMessage) SetFileData(v MessageFileData)`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *ChatapiMessage) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### GetVoiceData

`func (o *ChatapiMessage) GetVoiceData() MessageVoiceData`

GetVoiceData returns the VoiceData field if non-nil, zero value otherwise.

### GetVoiceDataOk

`func (o *ChatapiMessage) GetVoiceDataOk() (*MessageVoiceData, bool)`

GetVoiceDataOk returns a tuple with the VoiceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceData

`func (o *ChatapiMessage) SetVoiceData(v MessageVoiceData)`

SetVoiceData sets VoiceData field to given value.

### HasVoiceData

`func (o *ChatapiMessage) HasVoiceData() bool`

HasVoiceData returns a boolean if a field has been set.

### GetLocationData

`func (o *ChatapiMessage) GetLocationData() MessageLocationData`

GetLocationData returns the LocationData field if non-nil, zero value otherwise.

### GetLocationDataOk

`func (o *ChatapiMessage) GetLocationDataOk() (*MessageLocationData, bool)`

GetLocationDataOk returns a tuple with the LocationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationData

`func (o *ChatapiMessage) SetLocationData(v MessageLocationData)`

SetLocationData sets LocationData field to given value.

### HasLocationData

`func (o *ChatapiMessage) HasLocationData() bool`

HasLocationData returns a boolean if a field has been set.

### GetVideoData

`func (o *ChatapiMessage) GetVideoData() MessageVideoData`

GetVideoData returns the VideoData field if non-nil, zero value otherwise.

### GetVideoDataOk

`func (o *ChatapiMessage) GetVideoDataOk() (*MessageVideoData, bool)`

GetVideoDataOk returns a tuple with the VideoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoData

`func (o *ChatapiMessage) SetVideoData(v MessageVideoData)`

SetVideoData sets VideoData field to given value.

### HasVideoData

`func (o *ChatapiMessage) HasVideoData() bool`

HasVideoData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


