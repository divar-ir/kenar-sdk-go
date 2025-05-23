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

// checks if the ChatapiConversationSendMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatapiConversationSendMessageResponse{}

// ChatapiConversationSendMessageResponse struct for ChatapiConversationSendMessageResponse
type ChatapiConversationSendMessageResponse struct {
	Message *string `json:"message,omitempty"`
	Status *int32 `json:"status,omitempty"`
}

// NewChatapiConversationSendMessageResponse instantiates a new ChatapiConversationSendMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatapiConversationSendMessageResponse() *ChatapiConversationSendMessageResponse {
	this := ChatapiConversationSendMessageResponse{}
	return &this
}

// NewChatapiConversationSendMessageResponseWithDefaults instantiates a new ChatapiConversationSendMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatapiConversationSendMessageResponseWithDefaults() *ChatapiConversationSendMessageResponse {
	this := ChatapiConversationSendMessageResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ChatapiConversationSendMessageResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatapiConversationSendMessageResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ChatapiConversationSendMessageResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ChatapiConversationSendMessageResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChatapiConversationSendMessageResponse) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChatapiConversationSendMessageResponse) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChatapiConversationSendMessageResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ChatapiConversationSendMessageResponse) SetStatus(v int32) {
	o.Status = &v
}

func (o ChatapiConversationSendMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatapiConversationSendMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableChatapiConversationSendMessageResponse struct {
	value *ChatapiConversationSendMessageResponse
	isSet bool
}

func (v NullableChatapiConversationSendMessageResponse) Get() *ChatapiConversationSendMessageResponse {
	return v.value
}

func (v *NullableChatapiConversationSendMessageResponse) Set(val *ChatapiConversationSendMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChatapiConversationSendMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChatapiConversationSendMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatapiConversationSendMessageResponse(val *ChatapiConversationSendMessageResponse) *NullableChatapiConversationSendMessageResponse {
	return &NullableChatapiConversationSendMessageResponse{value: val, isSet: true}
}

func (v NullableChatapiConversationSendMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatapiConversationSendMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


