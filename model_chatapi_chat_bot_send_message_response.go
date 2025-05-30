/*
Kenar API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kenarapi

import (
	"encoding/json"
)

// checks if the ChatapiChatBotSendMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatapiChatBotSendMessageResponse{}

// ChatapiChatBotSendMessageResponse struct for ChatapiChatBotSendMessageResponse
type ChatapiChatBotSendMessageResponse struct {
	// ID of the conversation created or updated
	ConversationId *string `json:"conversation_id,omitempty"`
}

// NewChatapiChatBotSendMessageResponse instantiates a new ChatapiChatBotSendMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatapiChatBotSendMessageResponse() *ChatapiChatBotSendMessageResponse {
	this := ChatapiChatBotSendMessageResponse{}
	return &this
}

// NewChatapiChatBotSendMessageResponseWithDefaults instantiates a new ChatapiChatBotSendMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatapiChatBotSendMessageResponseWithDefaults() *ChatapiChatBotSendMessageResponse {
	this := ChatapiChatBotSendMessageResponse{}
	return &this
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *ChatapiChatBotSendMessageResponse) GetConversationId() string {
	if o == nil || IsNil(o.ConversationId) {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatapiChatBotSendMessageResponse) GetConversationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConversationId) {
		return nil, false
	}
	return o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *ChatapiChatBotSendMessageResponse) HasConversationId() bool {
	if o != nil && !IsNil(o.ConversationId) {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *ChatapiChatBotSendMessageResponse) SetConversationId(v string) {
	o.ConversationId = &v
}

func (o ChatapiChatBotSendMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatapiChatBotSendMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConversationId) {
		toSerialize["conversation_id"] = o.ConversationId
	}
	return toSerialize, nil
}

type NullableChatapiChatBotSendMessageResponse struct {
	value *ChatapiChatBotSendMessageResponse
	isSet bool
}

func (v NullableChatapiChatBotSendMessageResponse) Get() *ChatapiChatBotSendMessageResponse {
	return v.value
}

func (v *NullableChatapiChatBotSendMessageResponse) Set(val *ChatapiChatBotSendMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChatapiChatBotSendMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChatapiChatBotSendMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatapiChatBotSendMessageResponse(val *ChatapiChatBotSendMessageResponse) *NullableChatapiChatBotSendMessageResponse {
	return &NullableChatapiChatBotSendMessageResponse{value: val, isSet: true}
}

func (v NullableChatapiChatBotSendMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatapiChatBotSendMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


