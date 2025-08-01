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

// checks if the PostSubmitEmergencyResidencePostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSubmitEmergencyResidencePostResponse{}

// PostSubmitEmergencyResidencePostResponse struct for PostSubmitEmergencyResidencePostResponse
type PostSubmitEmergencyResidencePostResponse struct {
	PostToken *string `json:"post_token,omitempty"`
}

// NewPostSubmitEmergencyResidencePostResponse instantiates a new PostSubmitEmergencyResidencePostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSubmitEmergencyResidencePostResponse() *PostSubmitEmergencyResidencePostResponse {
	this := PostSubmitEmergencyResidencePostResponse{}
	return &this
}

// NewPostSubmitEmergencyResidencePostResponseWithDefaults instantiates a new PostSubmitEmergencyResidencePostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSubmitEmergencyResidencePostResponseWithDefaults() *PostSubmitEmergencyResidencePostResponse {
	this := PostSubmitEmergencyResidencePostResponse{}
	return &this
}

// GetPostToken returns the PostToken field value if set, zero value otherwise.
func (o *PostSubmitEmergencyResidencePostResponse) GetPostToken() string {
	if o == nil || IsNil(o.PostToken) {
		var ret string
		return ret
	}
	return *o.PostToken
}

// GetPostTokenOk returns a tuple with the PostToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSubmitEmergencyResidencePostResponse) GetPostTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PostToken) {
		return nil, false
	}
	return o.PostToken, true
}

// HasPostToken returns a boolean if a field has been set.
func (o *PostSubmitEmergencyResidencePostResponse) HasPostToken() bool {
	if o != nil && !IsNil(o.PostToken) {
		return true
	}

	return false
}

// SetPostToken gets a reference to the given string and assigns it to the PostToken field.
func (o *PostSubmitEmergencyResidencePostResponse) SetPostToken(v string) {
	o.PostToken = &v
}

func (o PostSubmitEmergencyResidencePostResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSubmitEmergencyResidencePostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PostToken) {
		toSerialize["post_token"] = o.PostToken
	}
	return toSerialize, nil
}

type NullablePostSubmitEmergencyResidencePostResponse struct {
	value *PostSubmitEmergencyResidencePostResponse
	isSet bool
}

func (v NullablePostSubmitEmergencyResidencePostResponse) Get() *PostSubmitEmergencyResidencePostResponse {
	return v.value
}

func (v *NullablePostSubmitEmergencyResidencePostResponse) Set(val *PostSubmitEmergencyResidencePostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSubmitEmergencyResidencePostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSubmitEmergencyResidencePostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSubmitEmergencyResidencePostResponse(val *PostSubmitEmergencyResidencePostResponse) *NullablePostSubmitEmergencyResidencePostResponse {
	return &NullablePostSubmitEmergencyResidencePostResponse{value: val, isSet: true}
}

func (v NullablePostSubmitEmergencyResidencePostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSubmitEmergencyResidencePostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


