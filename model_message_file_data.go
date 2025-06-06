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

// checks if the MessageFileData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageFileData{}

// MessageFileData struct for MessageFileData
type MessageFileData struct {
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
	SizeBytes *string `json:"size_bytes,omitempty"`
}

// NewMessageFileData instantiates a new MessageFileData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageFileData() *MessageFileData {
	this := MessageFileData{}
	return &this
}

// NewMessageFileDataWithDefaults instantiates a new MessageFileData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageFileDataWithDefaults() *MessageFileData {
	this := MessageFileData{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *MessageFileData) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageFileData) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *MessageFileData) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *MessageFileData) SetLink(v string) {
	o.Link = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MessageFileData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageFileData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MessageFileData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MessageFileData) SetName(v string) {
	o.Name = &v
}

// GetSizeBytes returns the SizeBytes field value if set, zero value otherwise.
func (o *MessageFileData) GetSizeBytes() string {
	if o == nil || IsNil(o.SizeBytes) {
		var ret string
		return ret
	}
	return *o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageFileData) GetSizeBytesOk() (*string, bool) {
	if o == nil || IsNil(o.SizeBytes) {
		return nil, false
	}
	return o.SizeBytes, true
}

// HasSizeBytes returns a boolean if a field has been set.
func (o *MessageFileData) HasSizeBytes() bool {
	if o != nil && !IsNil(o.SizeBytes) {
		return true
	}

	return false
}

// SetSizeBytes gets a reference to the given string and assigns it to the SizeBytes field.
func (o *MessageFileData) SetSizeBytes(v string) {
	o.SizeBytes = &v
}

func (o MessageFileData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageFileData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SizeBytes) {
		toSerialize["size_bytes"] = o.SizeBytes
	}
	return toSerialize, nil
}

type NullableMessageFileData struct {
	value *MessageFileData
	isSet bool
}

func (v NullableMessageFileData) Get() *MessageFileData {
	return v.value
}

func (v *NullableMessageFileData) Set(val *MessageFileData) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageFileData) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageFileData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageFileData(val *MessageFileData) *NullableMessageFileData {
	return &NullableMessageFileData{value: val, isSet: true}
}

func (v NullableMessageFileData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageFileData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


